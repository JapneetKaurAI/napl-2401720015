// package main

// import "fmt"

// //function to change the value of the pointer
// func ChangeValue(p *int) {
// 	*p = 100
// }

// //creating the structure
// type Student struct {
// 	name string
// 	age  int
// }

// //now writing the main function

// func main() {

// 	num := 10
// 	fmt.Println("Value of num:", num)
// 	fmt.Println("Address of num:", &num)
// 	fmt.Println("Value using pointer:", *(&num))

// 	fmt.Println("\nBefore function:", num)

// 	ChangeValue(&num)

// 	fmt.Println("After function:", num)

// 	student := new(Student)
// 	student.name = "Japneet"
// 	student.age = 20

// 	fmt.Println("\nStudent Details:")
// 	fmt.Println("Name:", student.name)
// 	fmt.Println("Age:", student.age)
// }

package main

import "fmt"

// Function to change the value using a pointer
func changeValue(p *int) {
	*p = 50
}

// Structure
type Student struct {
	name  string
	age   int
	marks float64
}

func main() {

	// Part 1: Referencing and Dereferencing

	num := 10

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value using pointer:", *(&num))

	// Part 2: Passing pointer to a function

	fmt.Println("\nBefore function:", num)

	changeValue(&num)

	fmt.Println("After function:", num)

	// Part 3: Creating a structure using new()

	student := new(Student)

	// Accessing and modifying structure fields
	student.name = "Japneet"
	student.age = 20
	student.marks = 85.5

	fmt.Println("\nStudent Details:")
	fmt.Println("Name:", student.name)
	fmt.Println("Age:", student.age)
	fmt.Println("Marks:", student.marks)
}
