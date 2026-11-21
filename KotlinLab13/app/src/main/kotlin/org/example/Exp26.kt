
package org.example

// 1. Default parameters
fun studentDetails(
    name: String = "Japneet",
    course: String = "Computer Science"
) {
    println("Name: $name")
    println("Course: $course")
}

// 2. Named arguments are demonstrated in main()

// 3. Vararg function
fun printNumbers(vararg numbers: Int) {
    println("Numbers:")
    for (number in numbers) {
        println(number)
    }
}

// 4. Higher-order function
fun calculate(
    numbers: List<Int>,
    operation: (Int) -> Int
) {
    val result = numbers.map(operation)
    println("Result: $result")
}

// Experiment 26
fun exp26() {

    println("===== DEFAULT PARAMETERS =====")

    // Calling without arguments
    studentDetails()

    // Calling with arguments
    studentDetails("Kanak", "BCA")

    println("\n===== NAMED ARGUMENTS =====")

    // Arguments passed out of order
    studentDetails(
        course = "Artificial Intelligence",
        name = "Japneet"
    )

    println("\n===== VARARG FUNCTION =====")

    printNumbers(10, 20, 30, 40, 50)

    println("\n===== HIGHER-ORDER FUNCTION =====")

    val numbers = listOf(1, 2, 3, 4, 5)

    calculate(numbers) { it * it }
}
