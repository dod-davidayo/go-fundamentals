package arrays

import "fmt"

func Demo() {
	numbers := [5]int{10, 20, 30, 40, 50}

	for _, number := range numbers {
		fmt.Println(number)
	}
}
