package main

import (
	"fmt"
	"os"
)

const filename = "profit_calculator.txt"

func main() {

	revenue, err := getUserInput("Enter the revenue:")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expenses, err := getUserInput("Enter the expenses:")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	taxRate, err := getUserInput("Enter the tax rate (as a percentage):")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate/100)
	writeToFile(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Ratio of Earnings before tax to Earnings after tax: %.3f\n", ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

func getUserInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getUserInput(prompt)
	}

	if input <= 0 {
		return 0, fmt.Errorf("input must be greater than zero")
	}

	return input, nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var tax = ebt * taxRate
	var profit = ebt - tax
	var ratio = ebt / profit

	return ebt, profit, ratio
}

func writeToFile(ebt, profit, ratio float64) {
	resultsStr := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile(filename, []byte(resultsStr), 0644)

}
