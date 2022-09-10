package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, city string) (bool, bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
	isValidCity := switchCity(city)
	return isValidName, isValidEmail, isValidTicketCount, isValidCity
}
