package main

import (
	"fmt"
	"max/bank/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalalnceFile = "balance.txt"
const defaultBalance = 1000.00

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalalnceFile, defaultBalance)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------------------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at", randomdata.PhoneNumber())

	for {

		presentOptions()
		var choice int
		fmt.Print("Please enter your choice (1-4): ")
		fmt.Scan(&choice)
		fmt.Print("\n")

		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Your balance is $", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Your deposit: $")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid deposit amount. Please enter a positive number.")
				continue
			}
			accountBalance += deposit
			fileops.WriteFloatToFile(accountBalalnceFile, accountBalance)
			fmt.Println("Your new balance is $", accountBalance)
		case 3:
			var withdrawl float64
			fmt.Print("Your withdrawl: $")
			fmt.Scan(&withdrawl)
			if withdrawl <= 0 {
				fmt.Println("Invalid withdrawl amount. Please enter a positive number.")
				continue
			}
			if withdrawl > accountBalance {
				fmt.Println("Insufficient funds. You cannot withdraw more than your account balance.")
				continue
			}
			accountBalance -= withdrawl
			fileops.WriteFloatToFile(accountBalalnceFile, accountBalance)
			fmt.Println("Your new balance is $", accountBalance)

		default:
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
			// break in Golang here only breaks the switch statement, not the loop
			// so we need to return to exit the program
		}

	}

}
