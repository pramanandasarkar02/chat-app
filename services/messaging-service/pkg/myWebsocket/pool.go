package myWebsocket

import (
	"chat-app/messaging-service/internal/models"
	"log"

	"github.com/gorilla/websocket"
)
// write upgrader



type Pool struct {
	Clients    map[*Client]bool
	Broadcast  chan models.Message
	Register   chan *Client
	Unregister chan *Client
}

type Client struct {
	Conn *websocket.Conn
}

func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan models.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (c *Client) ReadMessages(pool *Pool) {
	defer func() {
		pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg models.Message
		if err := c.Conn.ReadJSON(&msg); err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		pool.Broadcast <- msg
	}
}