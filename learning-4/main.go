package main

import "fmt"

type Project struct {
	id    int
	title string
	price float64
}

func main() {
	hobies := [3]string{"Coding", "Gym", "Reading"}
	fmt.Println(hobies[0])
	fmt.Println(hobies[1:])

	hobiesSlice := []string{hobies[0], hobies[1]}
	fmt.Println(hobiesSlice)
	hobiesSlice = append(hobiesSlice, hobies[2])
	fmt.Println(hobiesSlice)

	hobiesSlice = hobies[1:]
	fmt.Println(hobiesSlice)

	goals := []string{"Learn Go", "Create a project"}
	goals[1] = "Learn Vue"
	goals = append(goals, "Create a good project")
	fmt.Println(goals)

	products := []Project{
		{id: 1, title: "Product 1", price: 10.0},
		{id: 2, title: "Product 2", price: 20.0},
		{id: 3, title: "Product 3", price: 30.0},
	}
	fmt.Println(products)
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
