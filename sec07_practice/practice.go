package main

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies = [3]string{"Reading", "Gaming", "Hiking"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	var slice1 = hobbies[0:2]
	fmt.Println(slice1)
	var slice2 = hobbies[:2]
	fmt.Println(slice2)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice1 = hobbies[1:3]
	fmt.Println(slice1)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	var courseGoals = []string{"Learn Go", "Build a project"}
	fmt.Println(courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Contribute to Open Source")
	fmt.Println(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		Title string
		ID    int
		Price float64
	}
	var products = []Product{
		{"Laptop", 1, 999.99},
		{"Smartphone", 2, 499.99},
	}
	products = append(products, Product{"Tablet", 3, 299.99})
	fmt.Println(products)
	// Output the products
	for _, product := range products {
		fmt.Printf("Title: %s, ID: %d, Price: %.2f\n", product.Title, product.ID, product.Price)
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
