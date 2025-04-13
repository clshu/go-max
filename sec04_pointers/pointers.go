package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age // Get the address of age

	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age pointer dereferenced:", *agePointer)

	fmt.Println("Age:", age)

	getAdultYears := getAdultYears(age)
	fmt.Println("Adult years:", getAdultYears)
	fmt.Println("Adult years pointer:", getAdultYearsPointer(&age))

	mutateAge(&age)
	fmt.Println("Mutated age:", age)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPointer(agePointer *int) int {
	return *agePointer - 18
}

func mutateAge(agePointer *int) {
	*agePointer = *agePointer - 18
}
