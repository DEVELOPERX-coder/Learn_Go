package main

import (
	"fmt"
	"time"
)

// Function to simulate work
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Simulate work taking different times
		time.Sleep(time.Millisecond * time.Duration(job*100))
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back
	}
}

func main() {
	fmt.Println("🚀 Parallel Processing Demo 🚀")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}
