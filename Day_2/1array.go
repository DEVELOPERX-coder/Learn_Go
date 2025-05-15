package main

import "fmt"

func main() {
	// Declare an array of 5 integers
	var scores [5]int

	// Assign values
	scores[0] = 95
	scores[1] = 89
	scores[2] = 70
	scores[3] = 82
	scores[4] = 93

	// Short declaration with values
	names := [3]string{"Alex", "Blake", "Casey"}

	fmt.Println("First score:", scores[0])
	fmt.Println("Second name:", names[1])
	fmt.Println("Array length:", len(scores))
}
