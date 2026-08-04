package slices

import "fmt"

// Write a function:

// func RemoveProduct(products []string, index int) []string

// Remove an element using slicing.

func RemoveProduct(products []string, index int) []string {

	// check if the index is valid
	if index < 0 || index >= len(products) {
		fmt.Println("Invalid index.")
		return products
	}

	// Remove the product using slicing
	products = append(products[:index], products[index+1:]...)

	fmt.Println("Product removed successfully!")

	return products
}

func TestRemoveProducts() {

	products := []string{"Laptop", "Smartphone", "Tablet", "Headphones", "Smartwatch"}

	fmt.Println("Products in the slice before: ", products)

	// Remove the product at index 2
	products = RemoveProduct(products, 2)

	fmt.Println("Products in the slice after: ", products)
}
