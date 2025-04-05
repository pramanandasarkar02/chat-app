package handlers

import (
	"chat-app/messaging-service/pkg/myWebsocket"
	"log"
	"net/http"
	// Local package for pool and client logic
	"github.com/gorilla/websocket"                      // External package for WebSocket functionality
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var pool = myWebsocket.NewPool()

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	client := &myWebsocket.Client{Conn: conn}
	pool.Register <- client
	go client.ReadMessages(pool)
}

func HandleMessages() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
		case message := <-pool.Broadcast:
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Printf("Error broadcasting: %v", err)
					client.Conn.Close()
					delete(pool.Clients, client)
				}
			}
		}
	}
}