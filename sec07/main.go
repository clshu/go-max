package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))
	userNames[0] = "Jack"
	userNames[1] = "Jill"
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 5)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.2
	courseRatings["C#"] = 4.3
	courseRatings["JavaScript"] = 4.6
	fmt.Println(courseRatings)

}
