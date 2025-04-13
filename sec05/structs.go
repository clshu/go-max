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
	outputUserDetails(appUser)
	outputUserDetails2(&appUser)

	var appUser2 user = user{}
	outputUserDetails(appUser2)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUserDetails(u user) {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At (Local):", u.createdAt.Local())
	fmt.Println("Created At (UTC):", u.createdAt.UTC())
	fmt.Println("Created At (Unix):", u.createdAt.Unix())
	fmt.Println("Created At (UnixNano):", u.createdAt.UnixNano())
	fmt.Println("Created At (UnixMicro):", u.createdAt.UnixMicro())
	fmt.Println("Created At (UnixMilli):", u.createdAt.UnixMilli())
	fmt.Println("Created At (ISO):", u.createdAt.Format(time.RFC3339))
	fmt.Println("Created At (RFC1123):", u.createdAt.Format(time.RFC1123))
	fmt.Println("Created At (RFC850):", u.createdAt.Format(time.RFC850))
	fmt.Println("Created At (ANSIC):", u.createdAt.Format(time.ANSIC))
	fmt.Println("Created At (Kitchen):", u.createdAt.Format(time.Kitchen))
	fmt.Println("Created At (Stamp):", u.createdAt.Format(time.Stamp))
	fmt.Println("Created At (StampMilli):", u.createdAt.Format(time.StampMilli))
	fmt.Println("Created At (StampMicro):", u.createdAt.Format(time.StampMicro))
	fmt.Println("Created At (StampNano):", u.createdAt.Format(time.StampNano))
	fmt.Println("----------------------------------------")
}

func outputUserDetails2(u *user) {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt.Local())
	fmt.Println("----------------------------------------")
}
