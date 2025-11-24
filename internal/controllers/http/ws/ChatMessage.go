package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/middleware"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mutex      sync.RWMutex
}

type Client struct {
	Hub     *Hub
	Conn    *websocket.Conn
	Send    chan []byte
	UserID  string
	GroupID string
}

type ChatHandler struct {
	Hub            *Hub
	MessageUsecase usecase.MessageUsecase
	UserUsecase    usecase.UserUsecase
}

type WebSocketMessage struct {
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	UserID    string      `json:"user_id,omitempty"`
	GroupID   string      `json:"group_id,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
}

type ChatMessage struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	GroupID   string `json:"group_id"`
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()
			message := WebSocketMessage{
				Type: "user_connected",
				Data: map[string]string{
					"user_id":  client.UserID,
					"group_id": client.GroupID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.BroadcastToGroup(data, client.GroupID)
			}
		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.Mutex.Unlock()
			message := WebSocketMessage{
				Type: "user_disconnected",
				Data: map[string]string{
					"user_id":  client.UserID,
					"group_id": client.GroupID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.BroadcastToGroup(data, client.GroupID)
			}
		case message := <-h.Broadcast:
			h.Mutex.RLock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			h.Mutex.RUnlock()
		}
	}
}

func (h *Hub) BroadcastToGroup(message []byte, groupID string) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	for client := range h.Clients {
		if client.GroupID == groupID {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

// HandleWebSocket handles WebSocket connections for chat.
// @Summary WebSocket Chat
// @Description Connects via WebSocket for direct and group chat messaging
// @Tags chat
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Success 101 {string} string "Switching Protocols"
// @Router /ws/chat [get]
func (ch *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok || userID == "" {
		http.Error(w, "user ID not found in context (middleware missing or invalid)", http.StatusInternalServerError)
		return
	}
	groupID, ok := r.Context().Value(middleware.GroupIDContextKey).(string)
	if !ok {
		groupID = ""
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		Hub:     ch.Hub,
		Conn:    conn,
		Send:    make(chan []byte, 256),
		UserID:  userID,
		GroupID: groupID,
	}

	client.Hub.Register <- client
	go client.WritePump()
	go client.ReadPump(ch)
}

func (c *Client) ReadPump(handler *ChatHandler) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
		var wsMessage WebSocketMessage
		if err := json.Unmarshal(messageBytes, &wsMessage); err != nil {
			log.Printf("Message unmarshal error: %v", err)
			continue
		}
		switch wsMessage.Type {
		case "chat_message":
			handler.HandleChatMessage(c, wsMessage)
		case "typing":
			handler.HandleTyping(c, wsMessage)
		case "join_group":
			handler.HandleJoinGroup(c, wsMessage)
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			return
		}
	}
	// Send close message when channel is closed
	c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func (ch *ChatHandler) HandleChatMessage(client *Client, wsMessage WebSocketMessage) {
	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}
	content, _ := data["content"].(string)
	recipientId, _ := data["recipientId"].(string)
	groupId, _ := data["groupId"].(string)
	messageType := wsMessage.Type
	message, err := entity.NewMessage(client.UserID, groupId, content)
	if err != nil {
		log.Printf("Failed to create message entity: %v", err)
		return
	}
	savedMessage, err := ch.MessageUsecase.CreateMessage(message)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
		return
	}
	user, err := ch.UserUsecase.GetUserByID(client.UserID)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return
	}
	var text string
	if savedMessage.Text != nil {
		text = *savedMessage.Text
	}
	chatMessage := ChatMessage{
		ID:        savedMessage.MessageID,
		Content:   text,
		UserID:    savedMessage.SenderID,
		GroupID:   savedMessage.GroupID,
		Username:  user.NickName,
		Timestamp: savedMessage.CreatedAt.Unix(),
	}
	response := WebSocketMessage{
		Type: "new_message",
		Data: chatMessage,
	}
	if responseData, err := json.Marshal(response); err == nil {
		if messageType == "direct" && recipientId != "" {
			ch.Hub.SendToUser(responseData, recipientId)
		} else if messageType == "group" && groupId != "" {
			ch.Hub.BroadcastToGroup(responseData, groupId)
		}
	}
}

func (h *Hub) SendToUser(message []byte, userID string) {
	h.Mutex.RLock()
	defer h.Mutex.RUnlock()
	for client := range h.Clients {
		if client.UserID == userID {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
}

func (ch *ChatHandler) HandleTyping(client *Client, wsMessage WebSocketMessage) {
	response := WebSocketMessage{
		Type:   "typing",
		UserID: client.UserID,
		Data: map[string]interface{}{
			"user_id":  client.UserID,
			"group_id": client.GroupID,
			"typing":   wsMessage.Data,
		},
	}
	if responseData, err := json.Marshal(response); err == nil {
		ch.Hub.BroadcastToGroup(responseData, client.GroupID)
	}
}

func (ch *ChatHandler) HandleJoinGroup(client *Client, wsMessage WebSocketMessage) {

	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}
	newGroupID, ok := data["group_id"].(string)
	if !ok {
		return
	}
	if client.GroupID != "" {
		leaveMessage := WebSocketMessage{
			Type: "user_left",
			Data: map[string]string{
				"user_id":  client.UserID,
				"group_id": client.GroupID,
			},
		}
		if data, err := json.Marshal(leaveMessage); err == nil {
			ch.Hub.BroadcastToGroup(data, client.GroupID)
		}
	}
	client.GroupID = newGroupID
	joinMessage := WebSocketMessage{
		Type: "user_joined",
		Data: map[string]string{
			"user_id":  client.UserID,
			"group_id": client.GroupID,
		},
	}
	if data, err := json.Marshal(joinMessage); err == nil {
		ch.Hub.BroadcastToGroup(data, client.GroupID)
	}

	// 新規グループ作成時にAIエージェントの自己紹介メッセージを送信
	if client.UserID != "ai_agent" && newGroupID != "" {
		aiMessage := WebSocketMessage{
			Type: "new_message",
			Data: map[string]interface{}{
				"id":        "ai_intro",
				"content":   "こんにちは!私はこのグループをサポートするAIです。皆さんが仲良くなれるようにお手伝いします。何か質問があれば気軽に話しかけてくださいね!",
				"user_id":   "ai_agent",
				"group_id":  newGroupID,
				"username":  "AIエージェント",
				"timestamp": 0,
			},
		}
		if data, err := json.Marshal(aiMessage); err == nil {
			ch.Hub.BroadcastToGroup(data, newGroupID)
		}
	}

}