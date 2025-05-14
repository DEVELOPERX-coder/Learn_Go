package main

import "fmt"

// Simple function with no parameters or return value
func sayHello() {
	fmt.Println("Hello, Go programmer!")
}

// Function with parameters
func greetPerson(name string) {
	fmt.Println("Hello,", name)
}

// Function with return value
func addNumbers(a int, b int) int {
	return a + b
}

// Function with multiple return values
func getStats(numbers []int) (int, int) {
	sum := 0
	max := numbers[0]

	for _, num := range numbers {
		sum += num
		if num > max {
			max = num
		}
	}

	return sum, max
}

func main() {
	// Call functions
	sayHello()

	greetPerson("Champion")

	result := addNumbers(5, 7)
	fmt.Println("5 + 7 =", result)

	// Using multiple return values
	scores := []int{85, 93, 76, 88}
	total, highest := getStats(scores)
	fmt.Println("Total:", total)
	fmt.Println("Highest score:", highest)
}
