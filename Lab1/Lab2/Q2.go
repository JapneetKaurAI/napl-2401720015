package main

import (
	"fmt"
	"slices"
)

func main() {

	// SLICES

	// Taking a slice of marks of students from the user
	var marks []int

	fmt.Print("Enter the number of students: ")
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var mark int

		fmt.Printf("Enter the marks of student %d: ", i+1)
		fmt.Scan(&mark)

		marks = append(marks, mark)
	}

	fmt.Println("Student marks:", marks)

	// Appending a new element
	marks = append(marks, 90)

	fmt.Println("Updated marks after adding a new element:", marks)

	// Removing an element from the slice by its index
	// Index 1 means the second element
	if len(marks) >= 2 {
		marks = slices.Delete(marks, 1, 2)
	}

	fmt.Println("Updated marks after removing element at index 1:", marks)

	// Updating an element
	// Index 3 means the fourth element
	if len(marks) > 3 {
		marks[3] = 77
		fmt.Println("Updated marks after changing the fourth mark to 77:", marks)
	}

	// MAPS

	// Declaring a map
	stud_marks := map[string]int{}

	fmt.Print("\nEnter number of students: ")
	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {

		var name string
		var mark int

		fmt.Printf("Enter the name of student %d: ", i+1)
		fmt.Scan(&name)

		fmt.Printf("Enter the marks of student %d: ", i+1)
		fmt.Scan(&mark)

		stud_marks[name] = mark
	}

	fmt.Println("Student marks:", stud_marks)

	// Inserting a new pair into the map
	stud_marks["Khushi"] = 90

	fmt.Println("Updated student marks after adding Khushi:", stud_marks)

	// Deleting a pair from the map
	delete(stud_marks, "Anshu")

	fmt.Println("Updated student marks after removing Anshu:", stud_marks)

	// MAP LOOKUP

	// Looking up an existing key
	if marks, ok := stud_marks["japneet"]; ok {
		fmt.Println("Marks of Japneet:", marks)
	} else {
		fmt.Println("Japneet is not present in the map")
	}

	// Looking up a key that may not exist
	v, ok := stud_marks["kanak"]

	fmt.Println("Marks of Kanak:", v)
	fmt.Println("Kanak exists:", ok)
}

marks := []float{87.5,91.0}