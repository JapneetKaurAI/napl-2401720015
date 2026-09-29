package main

import (
	"fmt"
)

//2 types of function imported and exported
//exported function -> can be used outside the function
// func Add(a, b int) int{  //Add (capital first letter)
// 	return a + b
// }
// Golanguage.Add(2,3)

// name of the package. name of the function

// func main() {
// 	var a int
// 	fmt.Print("enter number 1:")
// 	fmt.Scan(&a)
// 	var b int
// 	fmt.Print("enter number 2:")
// 	fmt.Scan(&b)
// 	sum := mathutil.Add(a, b)
// 	fmt.Printf("sum=%d", sum)

// }
//Arrays ,Slices and Maps in GO
//[3]int -> size is part of the type
//marks[0]->insdex starts at zero
//in contiguous we share boundaries
//default values
//int-o
//float-0.0
//string-""
//bool- false

// var marks [3]int
// marks[0] = 90
// marks[1] = 85

// nums := [3]int{1,2,3} //sort and notation
// fmt.Println(len(nums))

// create an array which takes input from your user with 5 makrks
func main() {
	var arr [5]int
	fmt.Print("enter 5 subjects marks:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

}
