package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func BookTickets(availableTickets uint, usrTckts uint, bookings []string, fName string, lName string) (uint, []string, []string) {
	availableTickets -= usrTckts
	bookings = append(bookings, fName+" "+lName)
	return availableTickets, bookings, extractBookedFnames(bookings)
}

func SendTicket(userData UserData, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%d tickets for %s %s", userData.UsrTckts, userData.FName, userData.LName)
	Success().Println("\n#######################")
	Success().Printf("Sending ticket:\n%v \nto email address %s\n", ticket, userData.Email)
	Success().Println("#######################")
	wg.Done()
}

func extractBookedFnames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
	return firstNames
}
