package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p Person) ReadInputs() Person {
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)
	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)
	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
	return p

}
func (p Person) Printdata() {
	fmt.Println("Name  :", p.Name)
	fmt.Println("Age   :", p.Age)
	fmt.Println("Job   :", p.Job)
	fmt.Printf("Salary : %f", p.Salary)
}

func main() {
	var p1 Person
	var p2 Person

	fmt.Println("Person 1")
	p1 = p1.ReadInputs()

	fmt.Println("Person 2")
	p2 = p2.ReadInputs()

	fmt.Println("Details")
	p1.Printdata()
	p2.Printdata()

}
