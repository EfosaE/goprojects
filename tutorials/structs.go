package main

import "fmt"

// Define a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// func main() {
// 	// Create an instance of the struct
// 	person1 := Person{
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Age:       30,
// 	}

// 	// Access struct fields
// 	fmt.Println("First Name:", person1.FirstName)
// 	fmt.Println("Last Name:", person1.LastName)
// 	fmt.Println("Age:", person1.Age)

// 	// Modify struct fields
// 	person1.Age = 31
// 	fmt.Println("Updated Age:", person1.Age)

// 	// Create another instance using a shorthand
// 	person2 := Person{"Jane", "Smith", 25}
// 	fmt.Println("Person 2:", person2)
// }

type Vertex struct {
	X, Y int
}

var (
	age = 42
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	vertexPointer  = &v1 // has type *Vertex
	agePointer = &age // has type *int
)

func main() {
	fmt.Println(v1, vertexPointer, v2, v3, agePointer)
}