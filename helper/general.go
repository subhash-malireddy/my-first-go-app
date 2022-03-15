package helper

import "fmt"

func Greeting(conferenceName string, availableTickets uint, totalConferenceTickets uint) {
	Greet().Printf("✨✨ Welcome to %s booking application ✨✨.\n", conferenceName)
	fmt.Printf("%d tickets out of %d tickets are available.\n", availableTickets, totalConferenceTickets)
	fmt.Println("Get your tickets for the conference here!")
}
