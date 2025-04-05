package main

import (
	"log"
	"net/http"

	"chat-app/user-service/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	http.HandleFunc("/users", handlers.UserHandler)

	log.Println("User Service starting on :8082")
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}