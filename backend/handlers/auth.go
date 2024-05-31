package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/johansortiz/car-reservation-app/services"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var creds services.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if services.Authenticate(creds) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	} else {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
	}
}
