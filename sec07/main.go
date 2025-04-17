package main

import (
	"fmt"
)

func main() {
	var productNames = [4]string{"A Book"}
	prices := [4]float64{1.99, 2.99, 3.99, 4.99}
	temperature := []float64{1.99, 2.99, 3.99, 4.99}

	productNames[2] = "A Carpet"
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[3])
	fmt.Println(prices[1:3])
	fmt.Println(prices[:3])
	partialPrices := prices[1:]
	fmt.Println(partialPrices)
	highlightedPrices := partialPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Printf("partialPrices length: %v capacity: %v\n", len(partialPrices), cap(partialPrices))
	fmt.Printf("highlightedPrices length: %v capacity: %v\n", len(highlightedPrices), cap(highlightedPrices))

	fmt.Println("Temperatures: ", temperature)
	temperature = append(temperature, 5.99)
	fmt.Println("Temperatures: ", temperature)
}
