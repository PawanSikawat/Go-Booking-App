package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Multiverse"

const conferenceTickets uint = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName   string
	lastName    string
	email       string
	tickets     uint
	city        string
	bookingTime string
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets, city := getUserInput()
		isValidName, isValidEmail, isValidTicketCount, isValidCity := validateUserInput(firstName, lastName, email, userTickets, city)

		if isValidName && isValidEmail && isValidTicketCount && isValidCity {

			bookTicket(userTickets, firstName, lastName, email, city)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email, city)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next milennial.")
				break
			}
		} else {
			printIncorrectInput(isValidName, isValidEmail, isValidTicketCount, isValidCity)
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to the %v.\nWe have total of %v tickets and %v are still available.\nBootstrap for the further journey\n", conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var city string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Print("Enter the city of conference: ")
	fmt.Scan(&city)

	return firstName, lastName, email, userTickets, city
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, city string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		tickets:     userTickets,
		city:        city,
		bookingTime: time.Now().Format(time.RFC1123),
	}

	bookings = append(bookings, userData)

	fmt.Println(bookings)
	fmt.Printf("Type: %T\n", bookings)
	fmt.Printf("Size: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, city, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string, city string) {
	time.Sleep(30 * time.Second)
	ticket := fmt.Sprintf("%v %v tickets for %v %v", userTickets, city, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n%v\nTo email address: %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}

func printIncorrectInput(isValidName bool, isValidEmail bool, isValidTicketCount bool, isValidCity bool) {
	if !isValidName {
		fmt.Println("First Name or Last Name you entered is too short")
	}
	if !isValidEmail {
		fmt.Println("Email address you entered doesn't contain @ sign")
	}
	if !isValidTicketCount {
		fmt.Println("Number of tickets you entered is invalid")
	}
	if !isValidCity {
		fmt.Println("City of conference you entered is invalid. Please choose among [ New York, London, Berlin, Bangalore ]")
	}
}

func switchCity(city string) bool {
	switch city {
	case "New York":
		fmt.Println("Hey there. See you in US")
	case "London", "Berlin":
		fmt.Println("Hola. See you in Europe")
	case "Bangalore":
		fmt.Println("Namaste. See you in India")
	default:
		fmt.Println("You sure you are booking the right conference?")
		return false
	}
	return true
}
