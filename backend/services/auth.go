package services

import (
	"log"

	"github.com/johansortiz/car-reservation-app/models"
	"github.com/johansortiz/car-reservation-app/repository"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authenticate(creds Credentials) bool {
	db := repository.GetDB()
	var user models.User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username=$1 AND password=$2", creds.Username, creds.Password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println("Error authenticating user:", err)
		return false
	}
	return true
}
