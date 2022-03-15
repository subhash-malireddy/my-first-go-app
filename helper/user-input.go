package helper

import (
	"fmt"
	"strings"
)

type UserData struct {
	FName    string
	LName    string
	Email    string
	UsrTckts int
}

var userInput UserData

func GetUserInput() UserData {
	getfName()
	getlName()
	getemail()
	getusrTckts()
	return userInput
}

func getfName() {
	Input().Print("First Name : ")
	fmt.Scan(&userInput.FName)
	vfName()
}
func getlName() {
	Input().Print("Last Name : ")
	fmt.Scan(&userInput.LName)
	vlName()
}
func getemail() {
	Input().Print("Email : ")
	fmt.Scan(&userInput.Email)
	vemail()
}
func getusrTckts() {
	Input().Print("How many tickets do you want to book? ")
	fmt.Scan(&userInput.UsrTckts)
	vusrTckts()
}

// functions for verifying the user input
func vfName() {
	if !(len(userInput.FName) >= 2) {
		fmt.Println("First Name should be atleast 2 characters long.")
		getfName()
	}
}

func vlName() {
	if !(len(userInput.LName) >= 2) {
		Error().Println("Last Name should be atleast 2 characters long.")
		getlName()
	}
}

func vemail() {
	if !(strings.Contains(userInput.Email, "@")) {
		Error().Println("Email should have '@' symbol.")
		getemail()
	}
}

func vusrTckts() {
	if !(userInput.UsrTckts > 0) {
		Error().Println("Please provide a valid number that is greater than 0.")
		getusrTckts()
	}
}
