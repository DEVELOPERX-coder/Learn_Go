package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Task) Display() {
	status := " "
	if t.Completed {
		status = "✓"
	}

	fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
}

func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func SaveTasks(tasks []Task, filename string) error {
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, tasksJSON, 0644)
}

func LoadTasks(filename string) ([]Task, error) {
	var tasks []Task
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func findNextID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func main() {
	fmt.Println("📝 MY TODO LIST APP 📝")
	const filename = "tasks.json"

	tasks, err := LoadTasks(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		tasks = []Task{}
	}

	scanner := bufio.NewScanner(os.Stdin)
	nextID := findNextID(tasks)

	// tasks = append(tasks, Task{
	// 	ID:        1,
	// 	Title:     "Lean Go Programming",
	// 	Completed: false,
	// 	CreatedAt: time.Now(),
	// })

	// tasks = append(tasks, NewTask(1, "Learn Go Programming"))

	// scanner := bufio.NewScanner(os.Stdin)
	// nextID := 2

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1: View all tasks")
		fmt.Println("2: Add new task")
		fmt.Println("3: Mark task as completed")
		fmt.Println("4: Delete task")
		fmt.Println("5: Exit")

		fmt.Print("\nEnter your choice (1-5): ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// View all tasks
			fmt.Println("\nYour tasks:")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet! Add some tasks first.")
			} else {
				for _, task := range tasks {
					task.Display()
				}
			}

		case "2":
			// Add new task
			fmt.Print("\nEnter task title: ")
			scanner.Scan()
			title := scanner.Text()

			if title != "" {
				newTask := NewTask(nextID, title)
				tasks = append(tasks, newTask)
				nextID++
				fmt.Println("✅ Task added successfully!")

				// Save tasks
				if err := SaveTasks(tasks, filename); err != nil {
					fmt.Println("Error saving tasks:", err)
				}
			} else {
				fmt.Println("❌ Task title cannot be empty!")
			}

		case "3":
			// Mark task as completed
			fmt.Print("\nEnter task ID to mark as completed: ")
			scanner.Scan()
			var id int
			fmt.Sscanf(scanner.Text(), "%d", &id)

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					found = true
					fmt.Println("✅ Task marked as completed!")

					// Save tasks
					if err := SaveTasks(tasks, filename); err != nil {
						fmt.Println("Error saving tasks:", err)
					}
					break
				}
			}

			if !found {
				fmt.Println("❌ Task not found!")
			}

		case "4":
			// Delete task
			fmt.Print("\nEnter task ID to delete: ")
			scanner.Scan()
			var id int
			fmt.Sscanf(scanner.Text(), "%d", &id)

			for i, task := range tasks {
				if task.ID == id {
					// Remove task by taking elements before and after it
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("✅ Task deleted successfully!")

					// Save tasks
					if err := SaveTasks(tasks, filename); err != nil {
						fmt.Println("Error saving tasks:", err)
					}
					break
				}
			}

		case "5":
			// Exit
			fmt.Println("\nThank you for using the TODO App! Goodbye!")
			return

		default:
			fmt.Println("\n⚠️ Invalid choice! Please try again.")
		}
	}
}
