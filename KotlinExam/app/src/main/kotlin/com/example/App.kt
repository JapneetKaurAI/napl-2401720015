
package com.example

fun main() {
    question1()
    question2()
}

fun question1() {
    val name = "Japneet"
    var age = 20

    println("Name: $name")
    println("Age: $age")

    age = 21
    println("Updated Age: $age")
}

fun question2() {
    var city: String? = "Delhi"

    println("City: $city")

    city = null
    println("Updated City: $city")
}
