package main

import "fmt"

func main() {
	score := 85

	// Basic if statement
	if score > 70 {
		fmt.Println("You passed! Great job!")
	}

	// If-else
	if score >= 90 {
		fmt.Println("A grade - Outstanding work!")
	} else if score >= 80 {
		fmt.Println("B grade - Good job!")
	} else if score >= 70 {
		fmt.Println("C grade - Satisfactory")
	} else {
		fmt.Println("You need to study more")
	}

	// Go has a special if statement with initialization
	if num := 9; num < 10 {
		fmt.Println("Number is less than 10")
	}
}
