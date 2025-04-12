package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)

}
