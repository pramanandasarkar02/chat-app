package main

import (
	"chat-app/messaging-service/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	http.HandleFunc("/ws", handlers.WebSocketHandler)
	go handlers.HandleMessages()

	log.Println("Messaging Service starting on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}