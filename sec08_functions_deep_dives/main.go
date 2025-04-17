package main

func main() {
	sum := sumup(100, 1, 2, 3, 4, 5, 100)
	println("Sum of numbers:", sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
