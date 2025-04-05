package handlers

import (
	"encoding/json"
	"net/http"
	
	"chat-app/user-service/internal/models"
)

var users = make(map[string]models.User)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(users)
	case "POST":
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		users[user.Username] = user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}