package main

import "fmt"

func main() {
	// Create a map of string to int
	studentScores := map[string]int{
		"Alice": 92,
		"Bob":   85,
		"Carol": 97,
	}

	// Add or update a value
	studentScores["Dave"] = 89

	// Check if a key exists
	score, exists := studentScores["Bob"]
	if exists {
		fmt.Println("Bob's score:", score)
	} else {
		fmt.Println("Bob not found")
	}

	// Delete a key
	delete(studentScores, "Alice")

	// Loop through a map
	for name, score := range studentScores {
		fmt.Printf("%s: %d\n", name, score)
	}
}
