package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalalnceFile = "balance.txt"
const defaultBalance = 1000.00

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------------------")
	}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Your new balance is $", accountBalance)

		default:
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
			// break in Golang here only breaks the switch statement, not the loop
			// so we need to return to exit the program
		}

	}

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalalnceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {

	data, err := os.ReadFile(accountBalalnceFile)
	if err != nil {
		return defaultBalance, errors.New("could not read balance file. setting to default balance: " + fmt.Sprint(defaultBalance))
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {

		return defaultBalance, errors.New("could not parse balance from file. setting to default balance " + fmt.Sprint(defaultBalance))
	}
	if balance < 0 {

		return defaultBalance, errors.New("invalid balance in file. setting to default balance: " + fmt.Sprint(defaultBalance))
	}
	return balance, nil
}
