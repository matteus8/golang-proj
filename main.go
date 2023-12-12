package main

import "fmt"

func main() {
	//const states that this variable can't change in another place in the code
	var conferenceName = "'Go-lang Conference'"
	//not const and only var... but... conferenceName := "'Go-lang Conference'" would work
	const conferenceTicketAmount uint = 50
	var remainingTickets uint = 50
	//var conferenceTicketPrice = 15

	fmt.Printf("\nWelcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v avaiable.\n", conferenceTicketAmount, remainingTickets)
	fmt.Printf("Get your tickets here.\n")

	//get users first name
	var firstName string
	fmt.Printf("Please enter your first name... \n")
	fmt.Scan(&firstName)

	//get users last name
	var lastName string
	fmt.Printf("Please enter your last name...\n")
	fmt.Scan(&lastName)

	//get users email
	var userEmail string
	fmt.Printf("Please enter your email..\n")
	fmt.Scan(&userEmail)

	//get number of tickets user wants
	var userTickets uint
	fmt.Printf("Please enter how many tickets you would like...")
	fmt.Scan(&userTickets)
	fmt.Printf("\n\n\n")

	//print statement for booking
	fmt.Printf("User %v, %v, booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("Thank you %v, %v, for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("We will send an email too %v for confirmation.\n\n", userEmail)

	//adjust ticket booking
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("There are %v remaining tickets for %v conference\n\n", remainingTickets, conferenceName)

	//create an array with the users who bought the ticket
	//arrays have fixed size
	//var allBookings = [50]string{}
	//allBookings[25] = firstName + " " + lastName
	//fmt.Printf("The whole array %v", allBookings)

	//create a slice (an abstraction of an array) more powerfull

	//define an empty array that only accepts strings
	var allBookings = []string{}

	//append users to array
	allBookings = append(allBookings, firstName+" "+lastName)

	fmt.Printf("This is the users in bookings: %v", allBookings)
}
