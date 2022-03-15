package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var conferenceName = "Go Conference"
	const totalConferenceTickets uint = 50
	var availableTickets uint = totalConferenceTickets
	var bookings []string

	helper.Greeting(conferenceName, availableTickets, totalConferenceTickets)

	var userChoice string
	helper.Question().Println("Do you want to book tickets for the conference?")
	helper.Input().Print("Press any key and enter to book OR N | no to exit the app : ")
	fmt.Scan(&userChoice)
	if userChoice != "no" && userChoice != "N" {

		fmt.Println("Please enter your details.")
		fmt.Println()

		//ticket booking
		userData := helper.GetUserInput()
		var firstNames = []string{}
		availableTickets, bookings, firstNames = helper.BookTickets(availableTickets, uint(userData.UsrTckts), bookings, userData.FName, userData.LName)

		wg.Add(1)
		go helper.SendTicket(userData, &wg)
		//booking stats
		helper.Info().Printf("%d tickets are booked under the name %s %s for %s.\nEmail confirmation will be received on %s. \n", userData.UsrTckts, userData.FName, userData.LName, conferenceName, userData.Email)
		fmt.Printf("Total available tickets for the %s are %d \n", conferenceName, availableTickets)
		fmt.Printf("Currently there is/are %d booking(s). Attendees : %v\n", len(bookings), firstNames)
		if availableTickets == 0 {
			helper.Info().Println("✨ Tickets sold out! ✨")
		}
	} else {
		helper.Greet().Println("Thank you for using our booking-app.")
	}
	wg.Wait()
}
