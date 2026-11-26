data class Student(val name: String, val percentage: Double)

fun main() {

    // 1. LIST
    println("----- LIST -----")

    // Read-only List
    val list = listOf(10, 20, 30)
    println("Read-only List: $list")
    println("Contains 20: ${list.contains(20)}")
    println("Get element: ${list[0]}")

    // Mutable List
    val mutableList = mutableListOf(10, 20, 30)
    mutableList.add(40)
    mutableList.remove(20)
    println("Mutable List: $mutableList")


    // 2. SET
    println("\n----- SET -----")

    // Read-only Set
    val set = setOf(1, 2, 3)
    println("Read-only Set: $set")
    println("Contains 2: ${set.contains(2)}")

    // Mutable Set
    val mutableSet = mutableSetOf(1, 2, 3)
    mutableSet.add(4)
    mutableSet.remove(2)
    println("Mutable Set: $mutableSet")


    // 3. MAP
    println("\n----- MAP -----")

    // Read-only Map
    val map = mapOf(1 to "Apple", 2 to "Banana")
    println("Read-only Map: $map")
    println("Get key 1: ${map[1]}")
    println("Contains key 2: ${map.containsKey(2)}")

    // Mutable Map
    val mutableMap = mutableMapOf(1 to "Apple", 2 to "Banana")
    mutableMap[3] = "Mango"
    mutableMap.remove(2)
    println("Mutable Map: $mutableMap")


    // 4. FUNCTIONAL TRANSFORMATIONS
    println("\n----- TRANSFORMATIONS -----")

    val numbers = listOf(1, 2, 3, 4, 5)

    // map: Multiply each number by 2
    val mapped = numbers.map { it * 2 }
    println("Map: $mapped")

    // filter: Select even numbers
    val filtered = numbers.filter { it % 2 == 0 }
    println("Filter: $filtered")

    // reduce: Calculate sum
    val reduced = numbers.reduce { total, num -> total + num }
    println("Reduce: $reduced")


    // 5. SORTING STUDENTS
    println("\n----- SORTING STUDENTS -----")

    val students = listOf(
        Student("Aman", 85.5),
        Student("Riya", 92.0),
        Student("Karan", 78.0)
    )

    val sortedStudents = students.sortedByDescending { it.percentage }

    sortedStudents.forEach {
        println("${it.name} - ${it.percentage}%")
    }


    // 6. GROUP BY
    println("\n----- GROUP BY -----")

    val grouped = students.groupBy {
        if (it.percentage >= 85) "Excellent" else "Good"
    }

    println(grouped)
}