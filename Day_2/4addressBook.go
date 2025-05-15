package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create our address book using a map
	contacts := map[string]string{
		"John":  "555-1234",
		"Jane":  "555-5678",
		"Emily": "555-8765",
	}

	fmt.Println("📱 MINI ADDRESS BOOK 📱")

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1: View all contacts")
		fmt.Println("2: Add a contact")
		fmt.Println("3: Look up a contact")
		fmt.Println("4: Delete a contact")
		fmt.Println("5: Exit")

		var choice int
		fmt.Print("\nEnter your choice (1-5): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// View all contacts
			fmt.Println("\n📋 All Contacts:")
			if len(contacts) == 0 {
				fmt.Println("Address book is empty!")
			} else {
				for name, phone := range contacts {
					fmt.Printf("%-10s: %s\n", name, phone)
				}
			}

		case 2:
			// Add a contact
			var name, phone string
			fmt.Print("\nEnter name: ")
			fmt.Scanln(&name)
			name = strings.TrimSpace(name)

			fmt.Print("Enter phone: ")
			fmt.Scanln(&phone)
			phone = strings.TrimSpace(phone)

			contacts[name] = phone
			fmt.Println("✅ Contact added!")

		case 3:
			// Look up a contact
			var searchName string
			fmt.Print("\nEnter name to search: ")
			fmt.Scanln(&searchName)
			searchName = strings.TrimSpace(searchName)

			if phone, exists := contacts[searchName]; exists {
				fmt.Printf("Found: %s's phone is %s\n", searchName, phone)
			} else {
				fmt.Println("❌ Contact not found!")
			}

		case 4:
			// Delete a contact
			var deleteName string
			fmt.Print("\nEnter name to delete: ")
			fmt.Scanln(&deleteName)
			deleteName = strings.TrimSpace(deleteName)

			if _, exists := contacts[deleteName]; exists {
				delete(contacts, deleteName)
				fmt.Println("✅ Contact deleted!")
			} else {
				fmt.Println("❌ Contact not found!")
			}

		case 5:
			// Exit
			fmt.Println("\nThank you for using the address book!")
			return

		default:
			fmt.Println("\n⚠️ Invalid choice! Please try again.")
		}
	}
}
