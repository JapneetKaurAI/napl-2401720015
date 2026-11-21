
package org.example

class Student(
    var rollNo: Int,
    var name: String,
    var course: String,
    var percentage: Double
) {

    // Function to accept and update student details
    fun acceptData() {

        println("\nEnter Student Details")

        print("Enter Roll Number: ")
        rollNo = readln().toInt()

        print("Enter Student Name: ")
        name = readln()

        print("Enter Course: ")
        course = readln()

        print("Enter Percentage: ")
        percentage = readln().toDouble()
    }

    // Function to display student details
    fun displayDetails() {

        println("\n===== STUDENT DETAILS =====")
        println("Roll Number : $rollNo")
        println("Name        : $name")
        println("Course      : $course")
        println("Percentage  : $percentage%")
        println("===========================")
    }
}

// Experiment 27
fun exp27() {

    // Creating two Student objects
    val student1 = Student(
        101, "Japneet", "BSc Computer Science", 90.5
    )

    val student2 = Student(
        102, "Kanak", "BCA", 88.5
    )

    println("STUDENT 1")
    student1.acceptData()
    student1.displayDetails()

    println("\nSTUDENT 2")
    student2.acceptData()
    student2.displayDetails()
}
