package services

import (
	"log"

	"github.com/johansortiz/car-reservation-app/repository"
)

func Register(creds Credentials) bool {
	db := repository.GetDB()
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", creds.Username, creds.Password)
	if err != nil {
		log.Println("Error registering user:", err)
		return false
	}
	return true
}
