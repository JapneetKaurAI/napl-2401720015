
open System
open Microsoft.Data.Sqlite

// 1. DOMAIN MODEL USING F# TYPES

type BookStatus =
    | Available
    | Borrowed

type Book = {
    Id: int
    Title: string
    Author: string
    Year: int
    Status: BookStatus
}

// Database connection
let connectionString = "Data Source=library.db"

// Convert status into string
let statusToString status =
    match status with
    | Available -> "Available"
    | Borrowed -> "Borrowed"

// Convert string into status
let stringToStatus status =
    match status with
    | "Available" -> Available
    | "Borrowed" -> Borrowed
    | _ -> Available


// 2. CREATE DATABASE TABLE

let createTable () = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        "CREATE TABLE IF NOT EXISTS Books (
            Id INTEGER PRIMARY KEY AUTOINCREMENT,
            Title TEXT NOT NULL,
            Author TEXT NOT NULL,
            Year INTEGER NOT NULL,
            Status TEXT NOT NULL
        )"

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask
    return ()
}


// 3. CREATE OPERATION - ADD BOOK

let addBook title author year status = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        "INSERT INTO Books (Title, Author, Year, Status)
         VALUES (@Title, @Author, @Year, @Status)"

    command.Parameters.AddWithValue("@Title", title) |> ignore
    command.Parameters.AddWithValue("@Author", author) |> ignore
    command.Parameters.AddWithValue("@Year", year) |> ignore
    command.Parameters.AddWithValue("@Status", statusToString status) |> ignore

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask

    printfn "Book added successfully!"
}


// 4. READ OPERATION - DISPLAY ALL BOOKS

let getBooks () = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()
    command.CommandText <- "SELECT * FROM Books"

    use! reader =
        command.ExecuteReaderAsync() |> Async.AwaitTask

    let mutable books = []

    while reader.Read() do
        let book = {
            Id = reader.GetInt32(0)
            Title = reader.GetString(1)
            Author = reader.GetString(2)
            Year = reader.GetInt32(3)
            Status = stringToStatus (reader.GetString(4))
        }

        books <- book :: books

    return List.rev books
}


// 5. UPDATE OPERATION

let updateBook book = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        "UPDATE Books
         SET Title = @Title,
             Author = @Author,
             Year = @Year,
             Status = @Status
         WHERE Id = @Id"

    command.Parameters.AddWithValue("@Id", book.Id) |> ignore
    command.Parameters.AddWithValue("@Title", book.Title) |> ignore
    command.Parameters.AddWithValue("@Author", book.Author) |> ignore
    command.Parameters.AddWithValue("@Year", book.Year) |> ignore
    command.Parameters.AddWithValue("@Status", statusToString book.Status) |> ignore

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask

    printfn "Book updated successfully!"
}


// 6. DELETE OPERATION

let deleteBook id = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <- "DELETE FROM Books WHERE Id = @Id"
    command.Parameters.AddWithValue("@Id", id) |> ignore

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask

    printfn "Book deleted successfully!"
}


// 7. HIGHER-ORDER FUNCTIONS AND PATTERN MATCHING

let processBooks books =

    // Filter available books
    let availableBooks =
        books
        |> List.filter (fun book ->
            match book.Status with
            | Available -> true
            | Borrowed -> false
        )

    // Map books into display messages
    let bookDetails =
        availableBooks
        |> List.map (fun book ->
            sprintf "%s by %s (%d)"
                book.Title book.Author book.Year
        )

    printfn "\nAvailable Books:"

    bookDetails
    |> List.iter (printfn "%s")


// 8. MAIN FUNCTION

[<EntryPoint>]
let main argv =

    Async.RunSynchronously (async {

        printfn "===== LIBRARY MANAGEMENT SYSTEM ====="

        // Create database
        do! createTable ()

        // Add books
        do! addBook "Python Basics" "John Smith" 2024 Available
        do! addBook "Learning FSharp" "David Lee" 2025 Available
        do! addBook "Database Systems" "Robert" 2023 Borrowed

        // Read all books
        let! books = getBooks ()

        printfn "\nAll Library Records:"

        books
        |> List.iter (fun book ->
            printfn "ID: %d | Title: %s | Author: %s | Year: %d | Status: %s"
                book.Id
                book.Title
                book.Author
                book.Year
                (statusToString book.Status)
        )

        // Update first book and delete second book
        match books with
        | first :: second :: _ ->
            let updatedBook = { first with Status = Borrowed }
            do! updateBook updatedBook
            do! deleteBook second.Id

        | first :: [] ->
            let updatedBook = { first with Status = Borrowed }
            do! updateBook updatedBook

        | [] ->
            printfn "No books available."

        // Read updated records
        let! updatedBooks = getBooks ()

        // Apply map, filter and pattern matching
        processBooks updatedBooks

        printfn "\nLibrary application completed successfully!"
    })

    0
