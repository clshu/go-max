package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("Enter the revenue:")
	fmt.Scan(&revenue)
	fmt.Println("Enter the expenses:")
	fmt.Scan(&expenses)
	fmt.Println("Enter the tax rate (as a percentage):")
	fmt.Scan(&taxRate)

	profit, netProfit, ratio := calculateValues(revenue, expenses, taxRate/100)
	fmt.Println("Earnings before tax:", profit)
	fmt.Println("Earnings after tax:", netProfit)
	fmt.Println("Ratio of Earnings before tax to Earnings after tax:", ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	var profit = revenue - expenses
	var tax = profit * taxRate
	var netProfit = profit - tax
	var ratio = profit / netProfit

	return profit, netProfit, ratio
}
