package helper

import "strings"

// ?Export a function or variable from one package so that it will be available for all packages
// *Make capitalize first letter
func ValidateUserInput(firstName string, lastName string, emailId string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// Validation
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(emailId, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
