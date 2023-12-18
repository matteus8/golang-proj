package main

import (
	"fmt"
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

	// Create a slice (an abstraction of an array) for storing bookings
	var allBookings = []string{}

	// Continue taking user information and booking tickets until the conference is fully booked
	for remainingTickets > 0 {
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
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, please try again.\n", remainingTickets)
			continue // Continue to the next iteration of the loop
		}

		// Print statements for booking
		fmt.Printf("User %v, %v, booked %v tickets\n", firstName, lastName, userTickets)
		fmt.Printf("Thank you %v, %v, for booking %v tickets\n", firstName, lastName, userTickets)
		fmt.Printf("We will send an email to %v for confirmation.\n\n", userEmail)

		// Adjust remainingTickets after booking
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("There are %v remaining tickets for the %v conference\n\n", remainingTickets, conferenceName)

		// Append user's full name to the slice
		allBookings = append(allBookings, fmt.Sprintf("%v %v - Email: %v, Tickets: %v", firstName, lastName, userEmail, userTickets))
	}

	// Print a message when the conference is fully booked
	fmt.Printf("Our conference is fully booked. Hopefully next year is better.\n\n")

	// Print the list of users
	fmt.Println("List of Users:")
	for _, user := range allBookings {
		fmt.Println(user)
	}
}
