package main

import "fmt"

func main() {
	type strArray []string
	userNames := make(strArray, 2, 5)

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)
	fmt.Println("length:", len(userNames))
	fmt.Println("capacity:", cap(userNames))
	userNames[0] = "Jack"

	fmt.Println(userNames)

	type floatMap map[string]float64

	courseRatings := make(floatMap, 5)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.2
	courseRatings["C#"] = 4.3
	courseRatings["JavaScript"] = 4.6
	fmt.Println(courseRatings)

	fmt.Println("-----------------------------------------")

	for index, value := range userNames {
		fmt.Println(index, value)
	}
	fmt.Println("-----------------------------------------")
	for key, value := range courseRatings {
		fmt.Println(key, value)
	}

}
