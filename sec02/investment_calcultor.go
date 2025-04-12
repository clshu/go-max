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

	outputText("Investment Amount: $")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Investment Duration (years): ")
	fmt.Scan(&years)
	outputText("Inflation Rate (%): ")
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

	formattedFutureValue := fmt.Sprintf("Future Value of Investment: $%.2f\n", futureValue)
	formattedInflationAdjustedValue := fmt.Sprintf("Inflation Adjusted Value: $%.2f\n", inflationAdjustedValue)
	outputText(formattedFutureValue)
	outputText(formattedInflationAdjustedValue)

}

func outputText(text string) {
	fmt.Print(text)
}
