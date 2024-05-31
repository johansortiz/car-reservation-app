package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/johansortiz/car-reservation-app/models"
	"github.com/johansortiz/car-reservation-app/services"
)

func ReservationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		reservations := services.GetReservations()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reservations)
	} else if r.Method == http.MethodPost {
		var reservation models.Reservation
		err := json.NewDecoder(r.Body).Decode(&reservation)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		services.CreateReservation(reservation)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	}
}
