package maps

import "fmt"

// Create a map storing student scores.

// Example:

// John → 90
// Mary → 85
// David → 72

// Implement:

// Add
// Read
// Update
// Delete

func MapAssignment() {
	studentScore := map[string]string{
		"John":      90,
		"Mary":      85,
		"David":     72,
	}


	// Read Display all students and their score
	fmt.Println("Student Scores: ")
	for name, score := range studentScore {
		fmt.Printf("%s -> %d\n", name, score)
	}
    
	//  ADD
	studentScore["Osiki"] = 55 // create a new student data
	fmt.Println("\nadded Osiki.")

	// Update a new student
	studentScore["Mary"] = 10   // updating
	fmt.Println("Updated Mary's score.")

	// DELETE: 
	delete(studentScore. "David")
	fmt.Println("Deleted David")

	// Display the updated map
	fmt.Println("\nUpdated Student Scores: ")
	for name, score := range studentScore {
		fmt.Printf("%s -> %d\n", name, score)
	}
}
