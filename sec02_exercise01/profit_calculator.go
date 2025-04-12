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
	taxRate = taxRate / 100 // Convert percentage to decimal
	profit := revenue - expenses
	tax := profit * taxRate
	netProfit := profit - tax
	fmt.Println("Earnings before tax:", profit)
	fmt.Println("Earnings after tax:", netProfit)
	fmt.Println("Ratio of Earnings before tax to Earnings after tax:", profit/netProfit)
}
