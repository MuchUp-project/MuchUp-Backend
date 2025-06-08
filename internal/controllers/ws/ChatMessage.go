// internal/controllers/ws/ChatMessage.go
package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/usecase"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.RWMutex
}

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	userID  string
	groupID string
}

type ChatHandler struct {
	hub            *Hub
	messageUsecase usecase.MessageUsecase
	userUsecase    usecase.UserUsecase
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

func NewChatHandler(messageUsecase usecase.MessageUsecase, userUsecase usecase.UserUsecase) *ChatHandler {
	hub := &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	handler := &ChatHandler{
		hub:            hub,
		messageUsecase: messageUsecase,
		userUsecase:    userUsecase,
	}

	go hub.run()
	return handler
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()

			// 接続通知
			message := WebSocketMessage{
				Type: "user_connected",
				Data: map[string]string{
					"user_id":  client.userID,
					"group_id": client.groupID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.broadcastToGroup(data, client.groupID)
			}

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mutex.Unlock()

			// 切断通知
			message := WebSocketMessage{
				Type: "user_disconnected",
				Data: map[string]string{
					"user_id":  client.userID,
					"group_id": client.groupID,
				},
			}
			if data, err := json.Marshal(message); err == nil {
				h.broadcastToGroup(data, client.groupID)
			}

		case message := <-h.broadcast:
			h.mutex.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

func (h *Hub) broadcastToGroup(message []byte, groupID string) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for client := range h.clients {
		if client.groupID == groupID {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

func (ch *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	groupID := r.URL.Query().Get("group_id")

	if userID == "" || groupID == "" {
		http.Error(w, "user_id and group_id are required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:     ch.hub,
		conn:    conn,
		send:    make(chan []byte, 256),
		userID:  userID,
		groupID: groupID,
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump(ch)
}

func (c *Client) readPump(handler *ChatHandler) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, messageBytes, err := c.conn.ReadMessage()
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
			handler.handleChatMessage(c, wsMessage)
		case "typing":
			handler.handleTyping(c, wsMessage)
		case "join_group":
			handler.handleJoinGroup(c, wsMessage)
		}
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
		}
	}
}

func (ch *ChatHandler) handleChatMessage(client *Client, wsMessage WebSocketMessage) {
	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}

	content, ok := data["content"].(string)
	if !ok || content == "" {
		return
	}

	// メッセージをデータベースに保存
	message, err := entity.NewMessage(client.userID, client.groupID, content)
	if err != nil {
		log.Printf("Failed to create message entity: %v", err)
		return
	}

	savedMessage, err := ch.messageUsecase.CreateMessage(message)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
		return
	}

	user, err := ch.userUsecase.GetUserByID(client.userID)
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
		ch.hub.broadcastToGroup(responseData, client.groupID)
	}
}

func (ch *ChatHandler) handleTyping(client *Client, wsMessage WebSocketMessage) {
	response := WebSocketMessage{
		Type:   "typing",
		UserID: client.userID,
		Data: map[string]interface{}{
			"user_id":  client.userID,
			"group_id": client.groupID,
			"typing":   wsMessage.Data,
		},
	}

	if responseData, err := json.Marshal(response); err == nil {
		ch.hub.broadcastToGroup(responseData, client.groupID)
	}
}

func (ch *ChatHandler) handleJoinGroup(client *Client, wsMessage WebSocketMessage) {
	data, ok := wsMessage.Data.(map[string]interface{})
	if !ok {
		return
	}

	newGroupID, ok := data["group_id"].(string)
	if !ok {
		return
	}

	if client.groupID != "" {
		leaveMessage := WebSocketMessage{
			Type: "user_left",
			Data: map[string]string{
				"user_id":  client.userID,
				"group_id": client.groupID,
			},
		}
		if data, err := json.Marshal(leaveMessage); err == nil {
			ch.hub.broadcastToGroup(data, client.groupID)
		}
	}

	client.groupID = newGroupID

	joinMessage := WebSocketMessage{
		Type: "user_joined",
		Data: map[string]string{
			"user_id":  client.userID,
			"group_id": client.groupID,
		},
	}
	if data, err := json.Marshal(joinMessage); err == nil {
		ch.hub.broadcastToGroup(data, client.groupID)
	}
}
