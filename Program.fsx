
open System 
open Microsoft.Data.Sqlite

// Student record
type Student = {
    Id: int
    Name: string
    Marks: int
}

// Student data
let students = [
    { Id = 1; Name = "Japneet"; Marks = 90 }
    { Id = 2; Name = "Kanak"; Marks = 88 }
    { Id = 3; Name = "Anshu"; Marks = 60 }
]

// Pattern matching
let getGrade marks =
    match marks with
    | m when m >= 80 -> "A"
    | m when m >= 60 -> "B"
    | m when m >= 40 -> "C"
    | _ -> "Fail"

// Database connection
let connectionString = "Data Source=students.db"

// Create database table
let createTable () = async {
    use con = new SqliteConnection(connectionString)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText =
        "CREATE TABLE IF NOT EXISTS Students
         (Id INTEGER PRIMARY KEY, Name TEXT, Marks INTEGER)"

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Table created"
}

// MAP: Add 1 bonus mark
let bonus = List.map (fun s -> s.Marks + 1) students

printfn "Bonus Marks:"
bonus |> List.iter (printfn "%d")

// FILTER: Select students with 40 or more marks
let passed = List.filter (fun marks -> marks >= 40) bonus

printfn "Passed Students:"
passed |> List.iter (printfn "%d")

// Display grades
printfn "\nStudent Grades:"
students |> List.iter (fun s ->
    printfn "%s : %s" s.Name (getGrade s.Marks))

// ADD student
let addStudent s = async {
    use con = new SqliteConnection(connectionString)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText =
        "INSERT OR REPLACE INTO Students VALUES (@id,@name,@marks)"

    cmd.Parameters.AddWithValue("@id", s.Id) |> ignore
    cmd.Parameters.AddWithValue("@name", s.Name) |> ignore
    cmd.Parameters.AddWithValue("@marks", s.Marks) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Student added"
}

// READ students
let readStudents () = async {
    use con = new SqliteConnection(connectionString)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText = "SELECT * FROM Students"

    use reader = cmd.ExecuteReader()

    printfn "\nStudent Records:"

    while reader.Read() do
        printfn "%d %s %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))
}

// UPDATE student marks
let updateStudent id marks = async {
    use con = new SqliteConnection(connectionString)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText = "UPDATE Students SET Marks=@marks WHERE Id=@id"

    cmd.Parameters.AddWithValue("@id", id) |> ignore
    cmd.Parameters.AddWithValue("@marks", marks) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Student updated"
}

// DELETE student
let deleteStudent id = async {
    use con = new SqliteConnection(connectionString)
    do! con.OpenAsync() |> Async.AwaitTask

    use cmd = con.CreateCommand()
    cmd.CommandText = "DELETE FROM Students WHERE Id=@id"

    cmd.Parameters.AddWithValue("@id", id) |> ignore

    let! _ = cmd.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Student deleted"
}

// Main execution
let main () = async {

    do! createTable()

    // Add all students
    for s in students do
        do! addStudent s

    // Read students
    do! readStudents()

    // Update Kanak's marks
    do! updateStudent 2 95

    // Delete Anshu
    do! deleteStudent 3

    // Final records
    printfn "\nFinal Student Records:"
    do! readStudents()
}

// Run program
Async.RunSynchronously (main())


//difference between val and var (concept taken from kotlin)
