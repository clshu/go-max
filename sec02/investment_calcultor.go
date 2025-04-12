package main

import (
	"fmt"
	"math"
)

func main() {
	var infaltionRate float64
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: $")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Investment Duration (years): ")
	fmt.Scan(&years)
	fmt.Print("Inflation Rate (%): ")
	fmt.Scan(&infaltionRate)
	if investmentAmount <= 0 {
		fmt.Println("Investment amount must be greater than 0.")
		return
	}
	if expectedReturnRate <= 0 {
		fmt.Println("Expected return rate must be greater than 0.")
		return
	}
	if years <= 0 {
		fmt.Println("Investment duration must be greater than 0.")
		return
	}
	if infaltionRate <= 0 {
		fmt.Println("Inflation rate must be greater than 0.")
		return
	}
	// Calculate future value of investment
	// and adjust for inflation
	// Future Value = P * (1 + r)^n
	// Inflation Adjusted Value = FV / (1 + i)^n
	// where P = principal amount, r = rate of return, n = number of years, i = inflation rate

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var inflationAdjustedValue = futureValue / math.Pow(1+infaltionRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Inflation Adjusted Value: $", inflationAdjustedValue)

}
