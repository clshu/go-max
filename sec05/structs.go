package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {

	ufirstName := getUserData("Please enter your first name: ")
	ulastName := getUserData("Please enter your last name: ")
	ubirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	appUser = user{
		ufirstName,
		ulastName,
		ubirthdate,
		time.Now(),
	}
	fmt.Println(appUser)
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)

	var appUser2 user = user{}
	fmt.Println(appUser2)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
