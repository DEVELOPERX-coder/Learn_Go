package main

import "fmt"

// Define a custom type
type Person struct {
	Name     string
	Age      int
	Location string
}

func main() {
	// Create a new Person
	bob := Person{
		Name:     "Bob",
		Age:      29,
		Location: "New York",
	}

	// Access fields with dot notation
	fmt.Println("Name:", bob.Name)
	fmt.Println("Age:", bob.Age)

	// Update a field
	bob.Age = 30
	fmt.Println("Updated age:", bob.Age)

	// Create with shorthand (order matters!)
	alice := Person{"Alice", 27, "San Francisco"}
	fmt.Printf("%+v\n", alice) // %+v shows field names too!
}
