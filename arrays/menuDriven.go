package arrays

import "fmt"

// Create a menu-driven Student Management System using an array.

// Implement:

// Create
// Read
// Update
// Delete
// Exit

func StudentManagement() {
	// Array to store five product names

	var students [5]string

	var choice int
	var index int
	var student string

	for {
		fmt.Println("\n ---- Student Management System ------")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Exit the program")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		// Create
		case 1:
			fmt.Println("\nEnter 5 student names: ")

			for i := 0; i < len(students); i++ {
				fmt.Printf("Student %d: ", i)
				fmt.Scan(&students[i])
			}
		// READ
		case 2:
			fmt.Println("\nproducts in the array:")

			for i, student := range students {
				fmt.Printf("Index %d : %s\n", i, student)
			}

		// UPDATE
		case 3:
			fmt.Print("Enter the index to update (0-4): ")
			fmt.Scan(&index)

			if index >= 0 && index < len(students) {
				fmt.Print("Enter the new student name: ")
				fmt.Scan(&student)

				students[index] = student

				fmt.Println("Student updated successfully!")
			} else {
				fmt.Println("Invalid index.")
			}

		// DELETE
		case 4:
			fmt.Print("Enter the index to delete (0-4): ")
			fmt.Scan(&index)

			if index >= 0 && index < len(students) {
				students[index] = ""
				fmt.Println("Student deleted successfully")
			} else {
				fmt.Println("Invalid index")
			}

		// Exit
		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
