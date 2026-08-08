package mulitpleReturnValue

import "fmt"

// Create:

// StudentResult()

// Return:

// total
// average
// grade

func StudentResult(scores []float64) (float64, float64, string) {

	// Calculate total
	var total float64

	for _, score := range scores {
		total += score
	}

	// Calculate average
	average := total / float64(len(scores))

	// Determine grade
	var grade string

	if average >= 70 {
		grade = "A"
	} else if average >= 60 {
		grade = "B"
	} else if average >= 50 {
		grade = "C"
	} else if average >= 45 {
		grade = "D"
	} else if average >= 40 {
		grade = "E"
	} else {
		grade = "F"
	}

	return total, average, grade
}

func StudentResultDemo() {

	scores := []float64{84, 75, 65, 70, 90}

	// Receive the three returned values
	total, average, grade := StudentResult(scores)

	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Grade: %s\n", grade)
}
