
#r "nuget: MongoDB.Driver"

open MongoDB.Bson
open MongoDB.Driver

// Connect to MongoDB
let client = new MongoClient("mongodb://localhost:27017")

// Select database and collection
let database = client.GetDatabase("StudentDB")
let collection =
    database.GetCollection<BsonDocument>("Students")

// Create filter for student ID
let filter =
    Builders<BsonDocument>.Filter.Eq("StudentId", 1)

// Function to display all students
let displayStudents() =
    let students =
        collection.Find(FilterDefinition<BsonDocument>.Empty).ToList()

    printfn "\nStudent Records:"

    if students.Count = 0 then
        printfn "No student records found."

    for student in students do
        printfn "%s" (student.ToJson())

// Remove previous test record
collection.DeleteOne(filter) |> ignore

// CREATE operation
printfn "\n--- CREATE ---"

let student = BsonDocument()
student.Add("StudentId", 1)
student.Add("Name", "Japneet")
student.Add("Age", 20)

collection.InsertOne(student)
displayStudents()

// READ operation
printfn "\n--- READ ---"

let result = collection.Find(filter).FirstOrDefault()

if result <> null then
    printfn "Student Found: %s" (result.ToJson())

// UPDATE operation
printfn "\n--- UPDATE ---"

let update =
    Builders<BsonDocument>.Update.Set("Age", 21)

collection.UpdateOne(filter, update) |> ignore
displayStudents()

// DELETE operation
printfn "\n--- DELETE ---"

collection.DeleteOne(filter) |> ignore
displayStudents()

printfn "\nCRUD operations completed successfully."
