package arrays

// Create an array of 10 integers.

// Store the numbers 1–10.
// Print each number.
// Print the total sum.
// Print the average.

import "fmt"

func Array() {
	var sumAverage [10]int

	sumAverage[0] = 1
	sumAverage[1] = 2
	sumAverage[2] = 3
	sumAverage[3] = 4
	sumAverage[4] = 5
	sumAverage[5] = 6
	sumAverage[6] = 7
	sumAverage[7] = 8
	sumAverage[8] = 9
	sumAverage[9] = 10

	// variable to store the total sum
	sum := 0

	// Print each numbers and calculate the sum
	fmt.Println("Numbers in the array: ")
	for i := 0; i < len(sumAverage); i++ {
		fmt.Println(sumAverage[i])
		sum += sumAverage[i]
	}

	// Calculate the average
	average := float64(sum) / float64(len(sumAverage))

	// Println the results
	fmt.Println("\nTotal Sum:", sum)
	fmt.Printf("Average: %.2f\n", average)

}
