package main

import (
	"chat-app/auth-service/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)

	log.Println("Auth Service starting on :8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}