package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/johansortiz/car-reservation-app/services"
)

func CarsHandler(w http.ResponseWriter, r *http.Request) {
	criteria := r.URL.Query().Get("criteria")
	cars := services.SearchCars(criteria)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}
