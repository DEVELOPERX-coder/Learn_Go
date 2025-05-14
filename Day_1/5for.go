package main

import "fmt"

func main() {
	// Standard for loop (like in other languages)
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// While-style loop
	count := 0
	for count < 3 {
		fmt.Println("Count is:", count)
		count++
	}

	// Infinite loop with break
	sum := 0
	for {
		sum += 2
		if sum > 10 {
			break // Exit the loop
		}
		fmt.Println("Sum is now:", sum)
	}
}
