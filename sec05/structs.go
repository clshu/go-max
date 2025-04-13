package main

import (
	"fmt"
	"time"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.outputUserDetails()
	appUser.outputUserDetails2()
	appUser.clearUserName()
	appUser.outputUserDetails2()

	var appUser2 user = user{}
	appUser2.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, fmt.Errorf("all fields are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u user) outputUserDetails() {
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

func (u *user) outputUserDetails2() {
	fmt.Println("User Details:")
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt.Local())
	fmt.Println("----------------------------------------")
}
