package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Conference represents the conference details
type Conference struct {
	Name           string
	TicketAmount   uint
	RemainingTickets uint
	Bookings       []string
}

var conference = Conference{
	Name:           "Go-lang Conference",
	TicketAmount:   50,
	RemainingTickets: 50,
	Bookings:       []string{},
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Set the path to the templates folder
	router.LoadHTMLGlob("templates/*")

	// Set the path to the static folder
	router.Static("/static", "./static")

	// Set up routes
	router.GET("/", showHomePage)
	router.POST("/book", bookTickets)

	// Start the server
	router.Run(":8080")
}

func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"ConferenceName":     conference.Name,
		"ConferenceCapacity": conference.TicketAmount,
		"RemainingTickets":   conference.RemainingTickets,
	})
}

func bookTickets(c *gin.Context) {
	// Parse form data
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	email := c.PostForm("email")
	ticketsStr := c.PostForm("tickets")

	// Convert tickets to uint
	tickets, err := strconv.ParseUint(ticketsStr, 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number of tickets")
		return
	}

	// Check if there are enough tickets available
	if uint(tickets) > conference.RemainingTickets {
		c.String(http.StatusBadRequest, "Not enough tickets available")
		return
	}

	// Book tickets
	bookingDetails := fmt.Sprintf("%v %v - Email: %v, Tickets: %v", firstName, lastName, email, tickets)
	conference.Bookings = append(conference.Bookings, bookingDetails)
	conference.RemainingTickets -= uint(tickets)

	// Display confirmation message
	c.HTML(http.StatusOK, "confirmation.html", gin.H{
		"FirstName":          firstName,
		"LastName":           lastName,
		"BookedTickets":      tickets,
		"RemainingTickets":   conference.RemainingTickets,
		"ConferenceName":     conference.Name,
		"BookingConfirmation": "Booking successful! We will send an email for confirmation.",
	})
}
