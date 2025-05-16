package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Global counter for visitors
var visitorCount = 0

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	visitorCount++
	fmt.Fprintf(w, "<h1>Welcome to Your Go Web Server!</h1>")
	fmt.Fprintf(w, "<p>You are visitor #%d</p>", visitorCount)
	fmt.Fprintf(w, "<p><a href='/about'>About</a> | <a href='/counter'>Counter</a></p>")
}

// Handler for the about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About This Server</h1>")
	fmt.Fprintf(w, "<p>This is a simple web server written in Go!</p>")
	fmt.Fprintf(w, "<p><a href='/'>Home</a></p>")
}

// Handler for the counter page
func counterHandler(w http.ResponseWriter, r *http.Request) {
	// Get the count parameter if it exists
	countStr := r.URL.Query().Get("count")
	count := 10 // Default

	if countStr != "" {
		var err error
		count, err = strconv.Atoi(countStr)
		if err != nil || count < 1 {
			count = 10
		}
	}

	fmt.Fprintf(w, "<h1>Counter</h1>")

	// Generate the counters
	fmt.Fprintf(w, "<ul>")
	for i := 1; i <= count; i++ {
		fmt.Fprintf(w, "<li>Counter: %d</li>", i)
	}
	fmt.Fprintf(w, "</ul>")

	fmt.Fprintf(w, "<p>Change count: </p>")
	fmt.Fprintf(w, "<a href='/counter?count=5'>5</a> | ")
	fmt.Fprintf(w, "<a href='/counter?count=10'>10</a> | ")
	fmt.Fprintf(w, "<a href='/counter?count=20'>20</a>")
	fmt.Fprintf(w, "<p><a href='/'>Home</a></p>")
}

func main() {
	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/counter", counterHandler)

	// Start the server
	fmt.Println("🌐 Starting web server on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
