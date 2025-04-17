package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformFunction(2)
	triple := createTransformFunction(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformFunction(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}
