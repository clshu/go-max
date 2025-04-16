package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addGeneric(1.6, 2.5))
}

func addGeneric[T int | float64](a, b T) T {
	return a + b
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	return nil
}
