
fun main() {

    // 1. Nullable and Non-nullable Variables
    println("----- NULLABLE AND NON-NULLABLE -----")

    var name: String = "Japneet"
    var city: String? = "Delhi"

    println("Name: $name")
    println("City: $city")

    // name = null  // Compile-time error
    city = null

    println("City after assigning null: $city")


    // 2. SAFE CALL OPERATOR (?.)
    println("\n----- SAFE CALL OPERATOR -----")

    var studentName: String? = "Japneet"

    println("Name length: ${studentName?.length}")

    studentName = null

    println("Name length: ${studentName?.length}")


    // 3. ELVIS OPERATOR (?:)
    println("\n----- ELVIS OPERATOR -----")

    var username: String? = null

    val result = username ?: "Guest"

    println("Username: $result")


    // 4. NOT NULL ASSERTION OPERATOR (!!)
    println("\n----- NOT NULL ASSERTION -----")

    var address: String? = null

    try {
        println(address!!.length)
    }
    catch (e: NullPointerException) {
        println("NullPointerException caught!")
    }

    println("Program continues successfully.")
}
