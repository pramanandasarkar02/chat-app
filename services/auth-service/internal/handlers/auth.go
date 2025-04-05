package handlers

import (
	"chat-app/auth-service/internal/models"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret-key") // Should be in .env in production

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Dummy auth check (replace with DB in production)
	if creds.Username == "user" && creds.Password == "pass" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": creds.Username,
		})
		tokenString, _ := token.SignedString(jwtSecret)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Save to DB in production
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}