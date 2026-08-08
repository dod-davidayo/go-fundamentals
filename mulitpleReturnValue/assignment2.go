package mulitpleReturnValue

import "fmt"

func FindMinMAx(numbers []int) (int, int) {

	// Start by assuming the first number is both the minimum and maximum
	min := numbers[0]
	max := numbers[0]

	// Check each number in the slice
	for _, num := range numbers {

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}
	// return the minimum and maximum values
	return min, max
}

func FindMinMaxDemo() {
	numbers := []int{5, 2, 9, 1, 7}

	min, max := FindMinMAx(numbers)

	fmt.Printf("Numbers: %v\n", numbers)

	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
}