package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Ask for name
	fmt.Print("What's your name? ")

	// Read the input until user presses Enter
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove newline character

	// Use the input
	fmt.Printf("Hello, %s! You're awesome at learning Go!\n", name)

	// Ask for a number
	fmt.Print("Enter your favorite number: ")
	var number int
	fmt.Scan(&number)

	// Do something with the number
	doubled := number * 2
	fmt.Printf("%d doubled is %d\n", number, doubled)
}
