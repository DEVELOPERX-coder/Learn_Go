package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number between 1-20
	secretNumber := rand.Intn(20) + 1

	fmt.Println("🎮 NUMBER GUESSING GAME 🎮")
	fmt.Println("I'm thinking of a number between 1 and 20")

	// Track number of guesses
	attempts := 0

	// Game loop
	for {
		var guess int
		fmt.Print("Your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Too low! Try higher.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try lower.")
		} else {
			// Correct guess!
			fmt.Printf("🎉 YOU WIN! The number was %d\n", secretNumber)
			fmt.Printf("You got it in %d attempts\n", attempts)
			break
		}
	}

	r := rand.New(rand.NewSource(10))
	fmt.Println(r.Intn(20)) // Instead of using the current time
	fmt.Println(r.Intn(20))
	fmt.Println(r.Intn(20))
	fmt.Println(r.Intn(20))
}
