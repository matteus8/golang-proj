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
	var userName string
	fmt.Println("Please enter your first name...")
	fmt.Scan(&userName)

	var userTickets int
	var userTickets int

	//ask the user for their name

	fmt.Printf("User %v, booked %v tickets", userName, userTickets)

}
