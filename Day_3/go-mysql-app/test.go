package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import but don't use directly
)

func main() {
	fmt.Println("Connecting to MySQL...")

	// Format: username:password@tcp(host:port)/database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // Close when function exits

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
	id INT AUTO_INCREMENT PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	completed BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		panic(err.Error())
	}

	result, err := db.Exec("INSERT INTO tasks (title) VALUE (?)", "Learn MySQL with GO")
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error)
	}
	fmt.Printf("Inserted task with ID: %d\n", id)
	fmt.Println("\nAll tasks:")
	rows, err := db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var completed bool

		err = rows.Scan(&id, &title, &completed)
		if err != nil {
			panic(err.Error())
		}

		status := " "
		if completed {
			status = "✓"
		}

		fmt.Printf("[%s] %d: %s\n", status, id, title)
	}

	// Update a task
	_, err = db.Exec("UPDATE tasks SET completed = TRUE WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nMarked task %d as completed\n", id)

	// Delete a task (commented out to preserve data)
	/*
	   _, err = db.Exec("DELETE FROM tasks WHERE id = ?", id)
	   if err != nil {
	       panic(err.Error())
	   }
	   fmt.Printf("Deleted task %d\n", id)
	*/
}
