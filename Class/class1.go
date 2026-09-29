// package main //package program entry package

// import "fmt" //library(standard) has to be instilled , fmt stands for formating

// func main() { //execution starts here
// 	fmt.Println("Hello, World!")
// }
//go language is case sensitive
//{ } groups the statements
//no semicolons -> added automatically
//braces always required
//unused import fails -> build stops
//gofmt decides layout -> one standard style

// package main

// import "fmt"

// func main() {
// 	var num1 = 2.1
// 	var num2 = 2.2
// 	fmt.Println("addition", num1+num2)
// 	fmt.Println("subtraction", num2-num1)
// 	fmt.Println("multiplication", num2*num1)

// }

// package main

// import "fmt"

// func main() {
// 	var name = "Japneet"
// 	var roll = 15

// 	fmt.Printf("%v %d", name, roll) // P in println is capital because it is imported from the package
// }

// package main

// import "fmt"

// func main() {
// 	name := "Japneet"
// 	age := 18
// 	rollno := 2401720015
// 	fmt.Printf("%v %d %d", name, age, rollno)
// }

// package main

// import "fmt"

// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("enter your name: ")
// 	fmt.Scan(&name)

// 	fmt.Print("enter your age: ")
// 	fmt.Scan(&age)

// 	fmt.Print("hello", name)
// 	fmt.Print(" your age is ", age)
// }

//Take input for two numbers and print their sum.

// package main

// import "fmt"

// func main() {
// 	var num1 int
// 	var num2 int

// 	fmt.Print("enter number 1 : ")
// 	fmt.Scan(&num1)

// 	fmt.Print("enter number 2 : ")
// 	fmt.Scan(&num2)

// fmt.Println("the sum of two numbers is ", num1+num2)
// fmt.Println("the subtraction of two numbers is ", num2-num1)
// fmt.Println("the multiplication of two numbers is ", num2*num1)
// fmt.Println("the division of two numbers is ", num1/num2)
// }

//if else with input of numbers

// package main

// import "fmt"

// func main() {
// 	var marks int
// 	fmt.Print("enter your marks: ")
// 	fmt.Scan(&marks)
// 	if marks > 90 {
// 		fmt.Print("grade a")

// 	} else if marks > 75 {
// 		fmt.Print("grade B")
// 	} else {
// 		fmt.Print("Pass")
// 	}

// }

// session - 5
// data types and varibales in go
// only double quotes allowed
// integer -> int, int8....int64
// float -> float32 , float64
// String -> UTF-8 text ("") unicode transformation format
// Boolean -> true or false
// a variables type is fixed when declared
// rune = int32

// package main

// import "fmt"

// func main() {
// 	var rollno float32
// 	var name string
// 	var age int
// 	var kanak bool
// 	fmt.Printf("%f", rollno)
// 	fmt.Printf("%v", name)
// 	fmt.Printf("%v", age)
// 	fmt.Printf("%v", kanak)
// }

// func main() {
// 	var a float64 = 7.9
// 	b := int(a)
// 	fmt.Print(b)

// }

//it does not rounds off

// func main() {
// 	var a int = 65
// 	b := string(a)
// 	fmt.Print(b)
// }

//it returns ascii value small digits does not return anything and greater values goes to ascii value

// func main() {
// 	var a int = 65
// 	b := strconv.Itoa(a)
// 	fmt.Print(b)
// }

// package main

// func main() {
// 	s := "hello" (if with umblauts)
// 	len(s) // 6 bits
// 	len([]rune(s)) // 5 bits
// }

//rune counts bytes used to store that character

// take 3 integer inputs and print them with printf
// package main

// import "fmt"

// func main() {
// 	a := 10
// 	b := 20
// 	c := 30
// 	total := a + b + c
// 	avg := float64(total) / 3
// 	fmt.Printf("%.2f\n", avg)
// }

//learn real life applications
// banking wrong type loses money
// sensors int8 saves memory at scale
// databases column widths must match
// multilingual apps runs count characters
//conversion is explicit
// for is the only loop keyword
//marks code
// package main

// import "fmt"

// func main() {
// 	var marks int
// 	fmt.Print("enter your marks: ")
// 	fmt.Scan(&marks)

// 	if marks >= 40 {
// 		fmt.Println("pass")
// 	} else {
// 		fmt.Println("Fail")
// 	}
// }
//m% -> address of value not the value
// a variable can be declared in the function only not outside
//kepps scope small
//variable stays local
//declared amd tested together

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b int
// 	var c int
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	fmt.Print("enter b: ")
// 	fmt.Scan(&b)
// 	fmt.Print("enter c: ")
// 	fmt.Scan(&c)
// 	var total int
// 	total = a + b + c

// 	if avg := total / 3; avg > 40 {
// 		fmt.Println("pass")

// 	} else {
// 		fmt.Println("Fail")
// 	}
// }

// switch statement
// switch day {
// case "Sat","Sun":
// 	fmt.Println("holiday")
// default:
// 	fmt.Println("working")

// package main

// import "fmt"

// func main() {
// 	for {
// 		var choice int

// 		fmt.Println("\nEnter your choice:")
// 		fmt.Println("1. Integer")
// 		fmt.Println("2. Float")
// 		fmt.Println("3. Exit")
// 		fmt.Print("Choice: ")
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			var a int
// 			fmt.Print("Enter 1 integer: ")
// 			fmt.Scan(&a)

// 			var b int
// 			fmt.Print("Enter 2 integer: ")
// 			fmt.Scan(&b)

// 			fmt.Println("The sum of two numbers is:", a+b)
// 			fmt.Println("The subtraction of two numbers is:", a-b)
// 			fmt.Println("The multiplication of two numbers is:", a*b)
// 			fmt.Println("The division of two numbers is:", a/b)

// 		} else if choice == 2 {

// 			var a float64
// 			fmt.Print("Enter 1  number: ")
// 			fmt.Scan(&a)

// 			var b float64
// 			fmt.Print("Enter 2  number: ")
// 			fmt.Scan(&b)

// 			fmt.Println("The sum of two numbers is:", a+b)
// 			fmt.Println("The subtraction of two numbers is:", a-b)
// 			fmt.Println("The division of two numbers is:", a*b)

// 		} else if choice == 3 {
// 			fmt.Println("Good Bye!")
// 			break

// 		} else {
// 			fmt.Println("Invalid choice!")
// 		}
// 	}
// }

//tagged version in go language
//if else statement can be switched by switch version

// package main

// import "fmt"

// func main() {
// 	var marks int
// 	fmt.Print("enter your marks: ")
// 	fmt.Scan(&marks)
// 	switch {
// 	case marks >= 59:
// 		fmt.Print("first Division")
// 	case marks >= 50:
// 		fmt.Print("Second division ")
// 	case marks >= 40:
// 		fmt.Print("third division")
// 	case marks <= 40:
// 		fmt.Print("fourth division")
// 	default:
// 		fmt.Print("enter valid input ")
// 	}
// }

//tagged version multiple conditions are known
//untagged specific version

//for-> Go's only loop
//counting-> for i :=0; i<5; i++ {} //classic 3-part form

//while style -> for i < 5 {} //condition only

//Forever for {} exit with break

//range
//break->leaving the loop
//continue -> skip to next round
//range -> walks a string or collection

// package main

// import "fmt"

// func main() {

// 	for i, ch := range "Héllo" { //unicode character
// 		if i == 0 {
// 			continue //i==0 : will skip the first character
// 		}
// 		fmt.Println(i, string(ch))
// 	}
// }
//do while can be used as infinite for loop with break statement

// switch uses breaks by default
// package main

// import "fmt"

// func main() {
// 	for i := 1; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// //functions and Methods in Go

// func add(a int, b int) int {  //name of the variable followed by its type

// 	return a + b
// }

// shared type -> func add(a ,b int) int  //write the type once

// no return-> func greet(name string) //condition only

// two returns-> func divide(a, b int) (int, bool) // Go allows several

// package main

// import "fmt"

// func divide(a, b int) (int, bool) {
// 	if b == 0 {
// 		return 0, false
// 	}
// 	return a / b, true
// }
// func main() {
// 	var a int
// 	var _ int
// fmt.Print("enter a: ")
// fmt.Scan(&a)
// 	fmt.Print("enter b:")
// 	fmt.Scan(&_)

// 	fmt.Print(divide(a, _))

// }

// variadic functions

// func sum(nums.....int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

// package main

// import "fmt"

// func sum(nums ...int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }
// func main() {
// 	var a int
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&a)
// 	var b int
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&b)
// 	var c int
// 	fmt.Print("enter a: ")
// 	fmt.Scan(&c)
// 	result := sum(a, b, c)
// 	fmt.Println("Sum:", result)

// }

// methods functions with a receiver
// package main

// import "fmt"

// type celsius float64

// func (c celsius) Tof() float64 { //to define the method c->an object celsius->instance
// 	return float64(c)*1.8 + 32
// }
// func main() {
// 	var temp celsius
// 	fmt.Print("enter temperature")
// 	fmt.Scan(&temp)
// 	fmt.Print(temp.Tof())
// }

//take 3 integer value define a function computes average return the average and status
// package main

// import "fmt"

// func average(sub1, sub2, sub3 int) (float64, bool) {
// 	avg := float64(sub1+sub2+sub3) / 3
// 	if avg > 40 {
// 		return avg, true
// 	}
// 	return avg, false
// }
// func main() {
// 	var sub1 int
// 	fmt.Print("enter subject 1 marks:")
// 	fmt.Scan(&sub1)
// 	var sub2 int
// 	fmt.Print("enter subject 2 marks:")
// 	fmt.Scan(&sub2)
// 	var sub3 int
// 	fmt.Print("enter subject 3 marks:")
// 	fmt.Scan(&sub3)
// 	fmt.Print(average(sub1, sub2, sub3))
// }

// real-life application
// reusable logic -> write once , call anywhere
// two-value return result and status together
// unit testing small functions re testable
// team code clear inputs and outputs
// functions keep large programs readable
//pass by refrenece ->by address (can change the value by going to the main memory)
//pass by value ->make a copy and sending it seperately
//what is variadic functions?
// we use ... takes any count

// func abd() (int , bool) {

// }
// func main() {

// gofmt
// gowait
// package main

// import (
// 	"GoLanguage/utility"
// 	"fmt"
// )

// func main() {
// 	sum := utility.Add(5, 6)
// 	fmt.Printf("Sum=%v", sum)

// }
// package main

// import "fmt"

// func main() {
// 	var arr [5]int
// 	fmt.Println("enter 5 subjects marks:")
// 	for i := 0; i < 5; i++ {
// 		fmt.Scan(&arr[i])
// 	}
// 	fmt.Println("Array", arr)
// 	fmt.Println(len(arr))

// }

//slice-> index+value
//map-> key+value
//map order is random

//over a slice
// package main
// import "fmt"
// ages := map[string]int{}
// ages["Khushi"] = 21
// ages["Kanak"] = 22
// ages["Anshu"] = 23
// fmt.Print(ages["Khushi"])
// fmt.Print(ages["Kanak"])
// fmt.Print(ages["Anshu"])
// for i,ages := range s{
// 	fmt.Println(i,ages)
// }

// over a map
// package main

// import "fmt"

// func main() {

// 	ages := map[string]int{}

// 	ages["Khushi"] = 21
// 	ages["Kanak"] = 22
// 	ages["Anshu"] = 23

// 	fmt.Println(ages["Khushi"])
// 	fmt.Println(ages["Kanak"])
// 	fmt.Println(ages["Anshu"])

// 	for k, age := range ages {
// 		fmt.Println(k, age)
// 	}
// }

// create a slice which will store the marks of the three students and find its average
// package main

// import "fmt"

// func main() {
// 	s := []int{}
// 	s = append(s, 90)
// 	s = append(s, 99)
// 	s = append(s, 92)
// 	sum := 0
// 	for _, s := range s {
// 		sum += s

// 	}
// 	average := sum / len(s)
// 	fmt.Println("marks", s)
// 	fmt.Println("average", average)
// }

//instant lookup we use map(key) in caches
//config file key,values
//buffers a fixed array of bytes
//slice is a view it odes not create a copy because if we make changes in the copied array value changes itself in the original array
//method is alwys defined with the receiver
//
// func( s Student) Grade() string { //student is the receiver
// 	if s.Marks >=60 {
// 		return "First"

// 	}
// 	return "Second"
// }

// //a method is always assosciated with the class
// //structs are copied

// func birthday(s Student) {
// 	s.Age++
// }
// birthday(asha)
// fmt.Println(asha.age)

// package main
// import "fmt"
// type Student struct {
// 	Name string
// 	Age int
// 	Marks float32
// }
// func birthday(s,birthday) {
// 	s.Age++
//     fmt.Printf("inside function: %d",s.Age)

// }

// func main() {
// 	khushi := Student{
// 		Name :"Khushi"
// 		Age : 20
// 		Marks :99.0
// 	}
// 	birthday(Khushi)
// 	fmt.Printf("age outside fun : %d",Khushi.Age)
// }

// x := 10
// p := &x
// fmt.Println(*p) //10
// fmt.Println(p)  //0xc00....

// 0-pointer of a pointer is nil

// x:=10
// p:= &x
// *p=20
// fmt.Println(x) //x

//dereferencing
//pointers with functions
//pointers with methods

//when we dont use * ->value receiver
//usnused memory is reclaimed directly by garbage collector
//Go decides where values Stack or heap

// func (s *Student)

// package main
// import "fmt"

// type Students struct {
// 	Name string
// 	Age int
// 	Marks float64
// }
// func bonus(s students,grace_marks float64) float64 {
// 	return s.Marks + grace_marks
// }
// func main() {
// 	student1 := Students{"A",20,90.5}
// 	student2 :=Students{"B",21,92.5}
// 	student3 := Students{"C",22,98.5}

// 	calss := []Students{
// 		{"Astha" ,21,87.5}
// 		{"Ravi",22,80.5}
// 		{"A",20,90.5}
// 		{"b",20,85.5}
// 	}
// 	for _,s :=range class {
// 		fmt.Println(s.Name, s.Marks)
// 	}
// }

// Important * point  stores the address -> deferencing
// m% pointer

//sequential

//how to achieve concurrency in go language -> we can do that by GoRoutines -> the "go" keyword

// package main

// func main() {
// 	go count("japneet")
// 	go count("kanak")
// }

// func count(thing string) {
// 	for i := 1; true; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

//while calling we have to use go keyword .
//main does not wait for goroutines to exit

//sync.WaitGroup //for making the main function to let all other go routines finish
//we can use sleep but is not efficient how much time of sleep is required is not known
// reliable -> waitgroup

// imp for exam
// var wg sync.WaitGroup
// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	work()
// }()
// wg.Wait() //block till done

// defer -> done runs even on return
// wait block until all is done
// done make it finished

//what will be worked first we use channel

//make([]int ,5,10)

//

// package main
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go count("japneet")
// 	go count("kanak")
// }

// package main

// import "fmt"

// func printName(name string) {
// 	fmt.Println(name)
// }

// func main() {
// 	printName("Japneet")
// 	go printName("Kanak")
// }

// adding 2 second timer
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printName(name string) {
// 	fmt.Println(name)
// }

// func main() {
// 	go printName("Japneet")
// 	go printName("Kanak")

// 	time.Sleep(2 * time.Second)
// }

// now adding waitgroup
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func PrintName(name string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(name)

// }
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go PrintName("khushi", &wg)
// 	go PrintName("kanak", &wg)
// 	wg.Wait()
// }
//bufferd 3-4 datapoints we can

//2 types of channels
//sender and receiver

// ch:= make(chan int)
// go func() {
// 	ch<-42
// }()

// v := <-ch

//the select statement
//patterns to manage go routines
//channels

//1st -> worker pool

package main
import (
	"fmt"
	"time"
)
func worker(id int, job chan<- int, result <- chan int ) {
	for job in range(jobs) {
		fmt.Printf("worker %d started Job %d", id, job)
		time.sleep(time.Millisecond*500) 
	    results <- job *2
	}


}
func main() {
	numJobs := 100
	//create channel sender and receiver //bufferred channel becuase 100 jobs
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//start worker
	for w:=1; w<=3; w++ {
		go worker(w, jobs, results)

	}
	for j:=1; j<= ; j++ {
		jobs<-j
	}
	close(jobs)

	for a:= 1; a<= ; a++ {
		fmt.Println("Result:", <-)
	}


}