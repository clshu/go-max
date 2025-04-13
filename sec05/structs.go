package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.OutputUserDetails()
	appUser.OutputUserDetails2()
	appUser.ClearUserName()
	appUser.OutputUserDetails2()

	adminUser, err := user.NewAdmin("test@com", "tesst1234")
	if err != nil {
		fmt.Println("Error creating admin user:", err)
		return
	}

	adminUser.OutputUserDetails2()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails2()

	var appUser2 = user.User{}
	appUser2.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
