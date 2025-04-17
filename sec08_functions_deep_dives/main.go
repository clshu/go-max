package main

import (
	"fmt"
)

func main() {
	fact := factorial(5)
	factNonRecustion := factorialNonRecursion(5)
	fmt.Println("Factorial of 5:", fact)
	fmt.Println("Factorial of 5 (non-recursion):", factNonRecustion)

	fabNumber := 10
	fab := fibonacci(fabNumber)
	fabNonRecursion := fibonacciNonRecursion(fabNumber)
	fmt.Println("Fibonacci of 10:", fab)
	fmt.Println("Fibonacci of 10 (non-recursion):", fabNonRecursion)

}

// factorial of 5: 5! = 5 * 4 * 3 * 2 * 1 = 120
func factorialNonRecursion(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// fibonacci of 10:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciNonRecursion(n int) int {
	if n < 0 {
		panic("Negative number input")
	}
	fib := make([]int, n+1)
	fib[0] = 0
	if n >= 1 {
		fib[1] = 1
	}
	if n >= 2 {
		for i := 2; i <= n; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	// The sequence is: fib
	// Return the nth Fibonacci number
	return fib[n]
}
