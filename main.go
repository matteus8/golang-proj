package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define conferenceName as a constant string
	const conferenceName = "Go-lang Conference"

	// Define conferenceTicketAmount as a constant unsigned integer
	const conferenceTicketAmount uint = 50

	// Initialize remainingTickets with the initial available tickets
	var remainingTickets uint = 50

	fmt.Printf("\nWelcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v available.\n", conferenceTicketAmount, remainingTickets)
	fmt.Printf("Get your tickets here.\n")

	// Get user's first name
	var firstName string
	fmt.Printf("Please enter your first name... \n")
	fmt.Scan(&firstName)

	// Get user's last name
	var lastName string
	fmt.Printf("Please enter your last name...\n")
	fmt.Scan(&lastName)

	// Get user's email
	var userEmail string
	fmt.Printf("Please enter your email..\n")
	fmt.Scan(&userEmail)

	// Get the number of tickets the user wants
	var userTickets uint
	fmt.Printf("Please enter how many tickets you would like...")
	fmt.Scan(&userTickets)
	fmt.Printf("\n\n\n")

	// Check if there are enough tickets available
	if userTickets >= remainingTickets {
		fmt.Printf("We only have %v tickets remaining, please try again.\n", remainingTickets)
		return // Use 'return' instead of 'continue' and 'break' to exit the program
	}

	// Print statements for booking
	fmt.Printf("User %v, %v, booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("Thank you %v, %v, for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("We will send an email to %v for confirmation.\n\n", userEmail)

	// Adjust remainingTickets after booking
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("There are %v remaining tickets for the %v conference\n\n", remainingTickets, conferenceName)

	// Create a slice (an abstraction of an array) for storing bookings

	// Define an empty slice that only accepts strings
	var allBookings = []string{}

	// Append user's full name to the slice
	allBookings = append(allBookings, firstName+" "+lastName)

	fmt.Printf("This is the full name in bookings: %v\n", allBookings)

	// Create a slice for storing first names
	firstNames := []string{}
	for _, booking := range allBookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("This is the first name in bookings: %v\n", firstNames)

	// If there are no remaining tickets, end the program
	if remainingTickets == 0 {
		fmt.Printf("Our conference is fully booked. Please come back for the next one.\n")
	}
}
