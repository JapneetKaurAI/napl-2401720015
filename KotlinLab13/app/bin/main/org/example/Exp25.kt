
package org.example

fun exp25() {

    // 1. Hello Kotlin
    println("Hello Kotlin!!!")

    // 2. Employee Bonus Calculation
    print("Enter department code (IT/HR/FINANCE): ")
    val deptCode = readln().uppercase()

    print("Enter years of experience: ")
    val experience = readln().toInt()

    print("Enter employee salary: ")
    val salary = readln().toDouble()

    val bonusPercentage = when {
        deptCode == "IT" && experience > 5 -> 20
        deptCode == "HR" && experience > 3 -> 15
        deptCode == "FINANCE" && experience > 5 -> 18
        else -> 10
    }

    val bonusAmount = salary * bonusPercentage / 100

    println("Bonus Percentage: $bonusPercentage%")
    println("Bonus Amount: Rs. $bonusAmount")

    // 3. Average of individual digits
    print("Enter an integer: ")
    val number = readln().toLong()

    val digits = number.toString()
        .filter { it.isDigit() }
        .map { it.digitToInt() }

    val average = digits.average()

    println("Average of digits: $average")
}
