package main

import (
	"fmt"
	"time"
)

func square(num int, ch chan<- int) {
	fmt.Println("square routine has started")

	result := num * num
	ch <- result

	fmt.Println("square routine has completed")
}

func cube(num int, ch chan<- int) {
	fmt.Println("cube routine has started")

	result := num * num * num
	ch <- result

	fmt.Println("cube routine has completed")
}

func fibonacci(num int, ch chan<- int) {
	fmt.Println("fibonacci routine has started")

	a, b := 0, 1

	for i := 0; i < num; i++ {
		a, b = b, a+b
	}

	ch <- a

	fmt.Println("fibonacci routine has completed")
}

func main() {
	var num int

	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	// Creating a channel
	result := make(chan int)

	// Starting the goroutines
	go square(num, result)
	go cube(num, result)
	go fibonacci(num, result)

	for i := 1; i <= 3; i++ {
		fmt.Println(<-result)
		time.Sleep(time.Millisecond * 500)
	}

	// Receiving exactly 3 results
	// result1 := <-result
	// result2 := <-result
	// result3 := <-result

	// // Printing results
	// fmt.Println("Result 1 =", result1)
	// fmt.Println("Result 2 =", result2)
	// fmt.Println("Result 3 =", result3)
}
