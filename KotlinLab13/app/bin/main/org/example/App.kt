
package org.example

fun main() {

    println("========== KOTLIN LAB 13 ==========")
    println("1. Experiment 25 - Kotlin Basics")
    println("2. Experiment 26 - Function Declarations")
    println("3. Experiment 27 - Student Class")

    print("\nEnter your choice: ")

    when (readln().toInt()) {

        1 -> exp25()

        2 -> exp26()

        3 -> exp27()

        else -> println("Invalid choice!")
    }
}
