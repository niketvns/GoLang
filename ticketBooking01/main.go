package main

import (
	"fmt"
	"sync"
	"ticketBooking01/helper"
	"time"
)

// We have defined a custom data type
type UserData struct {
	firstName   string
	lastName    string
	emailId     string
	noOfTickets uint
}

// Package Level Variables
// * Best Practice: Define Variable as "Local" as Possible, create the variable where you need.
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, emailId, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailId, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, emailId)

		fmt.Println("Sending Your Ticket...")

		// To achieve concurrency/async in golang, we are using "go" keyword
		// Here sendTicket will run in background or in another thread that will not block my maink thread
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailId)

		var firstNames []string = getUsersFirstNames()

		fmt.Printf("The first name of bookings are: %v\n", firstNames)

		// var isNoTicketRemaining bool = remainingTickets == 0
		// isNoTicketRemaining := remainingTickets == 0

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference tickets are booked out, Come back next year.")
			// break
		} else {
			fmt.Printf("\n---------------------\n")
			fmt.Printf("Book Another Ticket:)")
			fmt.Printf("\n---------------------\n")
			fmt.Printf("%v tickets remaining for '%v'\n\n", remainingTickets, conferenceName)
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't container @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		fmt.Println("Please Try Again...")
	}
	// }

	wg.Wait()
}

// function declaration
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available to book\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend:)")
}

func getUsersFirstNames() []string {
	firstNames := []string{}
	// in go, "_" are used to identify unused variables, here "_" means index
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your Email Id: ")
	fmt.Scan(&emailId)
	fmt.Println("Enter No of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailId, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailId string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName

	// Custom UserData type
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		emailId:     emailId,
		noOfTickets: userTickets,
	}

	// *Create a map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailId"] = emailId
	// userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// push the element on the last index + 1 of slice
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive confirmation email at %v soon.\n", firstName, lastName, userTickets, emailId)
}

func sendTicket(userTickets uint, firstNme string, lastName string, emailId string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstNme, lastName)
	fmt.Printf("\n###############\n")
	fmt.Printf("Ticket Sent ✅\n%v to email address %v\n", ticket, emailId)
	fmt.Printf("###############\n\n")
	wg.Done()
}
