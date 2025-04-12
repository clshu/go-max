package main

import (
	"fmt"
	"math"
)

func main() {
	const infaltionRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Scan(&investmentAmount)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var inflationAdjustedValue = futureValue / math.Pow(1+infaltionRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Inflation Adjusted Value: $", inflationAdjustedValue)

}
