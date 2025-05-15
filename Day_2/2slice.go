package main

import "fmt"

func main() {
	// Create a slice
	scores := []int{95, 89, 70, 82, 93}

	// Add an element to the end
	scores = append(scores, 77)

	// Slice operations (start:end)
	firstTwo := scores[0:2] // Gets elements 0 and 1
	lastThree := scores[3:] // Gets elements from 3 onward

	fmt.Println("All scores:", scores)
	fmt.Println("First two:", firstTwo)
	fmt.Println("Last three:", lastThree)
	fmt.Println("Slice length:", len(scores))
	fmt.Println("Slice capacity:", cap(scores))
}
