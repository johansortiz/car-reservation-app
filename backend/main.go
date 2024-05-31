package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/johansortiz/car-reservation-app/handlers"
	"github.com/johansortiz/car-reservation-app/repository"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	repository.InitDB(dataSourceName)

	r := mux.NewRouter()

	r.HandleFunc("/api/auth", handlers.AuthHandler).Methods("POST")
	r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/cars", handlers.CarsHandler).Methods("GET")
	r.HandleFunc("/api/reservations", handlers.ReservationsHandler).Methods("GET", "POST")

	// Servir archivos estáticos (frontend)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
