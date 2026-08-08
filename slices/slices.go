package slices

import "fmt"

// Create a slice of employee names.

// Implement:

// Add Employee
// Display Employees
// Remove Employee by Index
// Exit

func Slices() {

	var AddEmployees []string

	var choice string
	var index int
	var AddEmployee string

	for {
		fmt.Println("\n --- Employee Management--- ")
		fmt.Println("A. POST")
		fmt.Println("B. GET")
		fmt.Println("C. PUT")
		fmt.Println("D. DELETE")
		fmt.Println("E. Exit the program")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// CREATE
		case "A", "a":
			var number int

			fmt.Print("How many employees do you want to add? ")
			fmt.Scan(&number)

			for i := 0; i < number; i++ {
				var employee string
				fmt.Printf("Enter employee %d name: ", i+1)
				fmt.Scan(&employee)
				AddEmployees = append(AddEmployees, employee)
			}

			// READ
		case "B", "b":
			fmt.Println("\nEmployee in the array")

			for i, AddEmployee := range AddEmployees {
				fmt.Printf("index %d : %s\n", i, AddEmployee)
			}

		// UPDATE
		case "C", "c":
			fmt.Print("Enter the index to update (0-4): ")
			fmt.Scan(&index)

			if index >= 0 && index < len(AddEmployees) {
				fmt.Print("Enter the AddEmployee Name: ")
				fmt.Scan(&AddEmployee)

				AddEmployees[index] = AddEmployee

				fmt.Println("Employee updated successfully!")
			} else {
				fmt.Println("Invalid index.")
			}

		// DELETE
		case "D", "d":
			fmt.Print("Enter the index to delete (0-4): ")
			fmt.Scan(&index)

			if index >= 0 && index < len(AddEmployees) {
				AddEmployees = append(AddEmployees[:index], AddEmployees[index+1:]...)
				fmt.Println("Employee deleted successfully!")
			} else {
				fmt.Println("Invalid index.")
			}
			// Exit:
		case "E", "e":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
