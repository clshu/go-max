package main

import (
	"fmt"
	"math"
)

const infaltionRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment Amount: $")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Investment Duration (years): ")
	fmt.Scan(&years)

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

	futureValue, inflationAdjustedValue := calculateValues(investmentAmount, expectedReturnRate, years)

	formattedFutureValue := fmt.Sprintf("Future Value of Investment: $%.2f\n", futureValue)
	formattedInflationAdjustedValue := fmt.Sprintf("Inflation Adjusted Value: $%.2f\n", inflationAdjustedValue)
	outputText(formattedFutureValue)
	outputText(formattedInflationAdjustedValue)

}

func outputText(text string) {
	fmt.Print(text)
}

// Calculate future value of investment
// and adjust for inflation
// Future Value = P * (1 + r)^n
// Inflation Adjusted Value = FV / (1 + i)^n
// where P = principal amount, r = rate of return, n = number of years, i = inflation rate

func calculateValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var inflationAdjustedValue = futureValue / math.Pow(1+infaltionRate/100, years)

	return futureValue, inflationAdjustedValue
}
