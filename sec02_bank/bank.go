package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("\nWhat would you like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please enter your choice (1-4): ")
		fmt.Scan(&choice)
		fmt.Print("\n")

		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
			continue
		}

		checkBalance := choice == 1
		depositMoney := choice == 2
		withdrawMoney := choice == 3

		if checkBalance {
			fmt.Println("Your balance is $", accountBalance)
		} else if depositMoney {
			var deposit float64
			fmt.Print("Your depost:$ ")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid deposit amount. Please enter a positive number.")
				continue
			}
			accountBalance += deposit
			fmt.Println("Your new balance is $", accountBalance)
		} else if withdrawMoney {
			var withdraw float64
			fmt.Print("Your withdraw:$ ")
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid withdraw amount. Please enter a positive number.")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Insufficient funds. You cannot withdraw more than your account balance.")
				continue
			}
			accountBalance -= withdraw
			fmt.Println("Your new balance is $", accountBalance)
		} else {

			break
		}

	}

	fmt.Println("Thank you for using Go Bank. Goodbye!")
}
