package main

func main() {
	numbers := []int{1, 2, 3, 4, 5, 100}
	sum := sumup(100, 1, 2, 3, 4, 5, 100)
	println("Sum of numbers:", sum)
	anotherSum := sumup(100, numbers...)
	println("Sum of another numbers:", anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
