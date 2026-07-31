package arrays

import "fmt"

// Create an array of 5 Book structs.

// Each book should contain:

// ID
// Title
// Author
// Price

// Implement CRUD operations.

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func CrudBook() {

	// Array to store 5 books
	var books [5]Book

	var choice int

	for {
		fmt.Println("\n --- Book Management System ----")
		fmt.Println("1. Create Book")
		fmt.Println("2. Read Book")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Exit the program")
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// CREATE
		case 1:
			fmt.Println("\nEnter details for 5 books")

			for i := 0; i < len(books); i++ {
				fmt.Printf("\nBook %d\n", i+1)

				fmt.Print("Enter ID: ")
				fmt.Scan(&books[i].ID)

				fmt.Print("Enter Title: ")
				fmt.Scan(&books[i].Author)

				fmt.Print("Enter Author: ")
				fmt.Scan(&books[i].Author)

				fmt.Print("Enter Price: ")
				fmt.Scan(&books[i].Price)
			}
			fmt.Println("\nBooks added successfully!")

		// READ
		case 2:
			fmt.Println("\nBook Records")

			for _, book := range books {
				// skip deleted books

				if book == (Book{}) {
					continue
				}

				fmt.Printf("ID: %d | Title: %s | Author: %s\n",
					book.ID,
					book.Title,
					book.Author)
			}
			// UPDATE
		case 3:
			var id int
			var found bool

			fmt.Print("Enter Book ID to update: ")
			fmt.Scan(&id)

			for i := 0; i < len(books); i++ {
				if books[i].ID == id {
					fmt.Print("Enter New Title: ")
					fmt.Scan("&books[i].Title")

					fmt.Print("Enter New Author: ")
					fmt.Scan(&books[i].Author)

					found = true
					fmt.Println("Book updated successfully!")
					break
				}
			}
			if !found {
				fmt.Println("Book not found.")
			}
		// DELETE
		case 4:
			var id int
			var found bool

			fmt.Println("Enter Book ID to delete: ")
			fmt.Scan(&id)

			for i := 0; i < len(books); i++ {
				if books[i].ID == id {
					books[i] = Book{} //Empty struct
					found = true
					fmt.Println("Book deleted sucessfully!")
					break
				}
			}

			if !found {
				fmt.Println("Book not found. ")
			}

		// Exit
		case 5:
			fmt.Println("Exited the program.")
			return

		default:
			fmt.Println("Invalid choice. ")
		}
	}
}
