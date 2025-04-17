package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(number *[]int, transform transformFunc) []int {
	dNumbers := make([]int, len(*number))
	for i, n := range *number {
		dNumbers[i] = transform(n)
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(n int) int {
	return n * 3
}
