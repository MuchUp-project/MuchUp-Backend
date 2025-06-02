package handlers 

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleChatMessage(w http.ResponseWriter,r *http.Request) {
	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
        log.Println(err)
        return
    }
	defer conn.Close()

	for {
		_,mes,err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Println("resv:%s",string(mes))

		conn.WriteMessage(websocket.TextMessage,mes)

	}
}