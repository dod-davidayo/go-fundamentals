package maps

import "fmt"


// Create a phonebook.

// Use:

// map[string]string

// where

// Name → Phone Number

// Implement CRUD.


func PhoneBook() {
	phoneBook := map[string]string{
		"David": "08036461531",
		"Ayo":   "08036461532",
		"Osiki": "08036461533",
	}

	// CREATE 
	phoneBook["Grace"] = "08036461534" // create a new contact	

	fmt.Println("After adding Grace: ")
	for name, number := range phoneBook {
		fmt.Printf("%s -> %s\n", name, number)
	}

	// READ
	fmt.Println("\nPhonebook: ")

	if number, found := phoneBook["Mary"]; found {
		fmt.Println(number)
	} else {
		fmt.Println("Mary not found in the phonebook.")
	}

	// UPDATE
	phoneBook["Ayo"] = "08036461535"

	fmt.Println("\nAfter updating Ayo's number: ")
	for name, number := range phoneBook {
		fmt.Printf("%s -> %s\n", name, number)
	}
	
	fmt.Println("Ayo's phone number updated successfully.")

	// DELETE
	delete(phoneBook, "David")
	fmt.Println("\nAfter deleting David: ")
	for name, number := range phoneBook {
		fmt.Printf("%s -> %s\n", name, number)
	}
	
	fmt.Println("David removed successfully.")

	// Display final book

	fmt.Println("\nFinal Phonebook: ")

	for name, number := range phoneBook {
		fmt.Printf("%s -> %s\n", name, number)
	}
}