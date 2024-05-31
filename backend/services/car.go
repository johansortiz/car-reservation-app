package services

import "github.com/johansortiz/car-reservation-app/models"

func SearchCars(criteria string) []models.Car {
	// Implement search logic
	cars := []models.Car{
		{ID: 1, Model: "Sedan", Details: "A nice sedan car"},
		{ID: 2, Model: "SUV", Details: "A powerful SUV"},
	}
	return cars
}
