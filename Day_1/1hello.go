package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// Declare some variables
	name := "Coding Champion"
	age := 25
	isLearning := true

	// Use our variables
	fmt.Println("Your name is:", name)
	fmt.Println("Your age is:", age)

	if isLearning {
		fmt.Println("You're doing great learning Go!")
	}

	fmt.Println("This is mix ", name, " ", age)
}
