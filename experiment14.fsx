
// Base class
type Shape() =
    abstract member Area : unit -> float
    default this.Area() = 0.0

// Derived class: Circle
type Circle(radius: float) =
    inherit Shape()

    override this.Area() =
        3.14 * radius * radius

// Derived class: Rectangle
type Rectangle(length: float, breadth: float) =
    inherit Shape()

    override this.Area() =
        length * breadth

// Main program
let shapes: Shape list =
    [
        Circle(5.0) :> Shape
        Rectangle(4.0, 6.0) :> Shape
    ]

// Demonstrating polymorphism
for shape in shapes do
    printfn "Area = %.2f" (shape.Area())
