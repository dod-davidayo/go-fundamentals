package slices

import "fmt"

// Create a shopping cart.

// Each product should have:

// Name
// Price

// Allow the user to:

// Add items
// Remove items
// View cart
// Calculate total price

type Product struct {
	Name  string
	Price float64
}

func ShoppingCart() {
	var carts []Product

	var choice string
	var index int
	//var cart string

	for {
		fmt.Println("\n --- Shopping Cart --- ")
		fmt.Println("A. Add Item")
		fmt.Println("B. Remove Item")
		fmt.Println("C. View Cart")
		fmt.Println("D. Calculate Total Price")
		fmt.Println("E. Exit the program")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		// CREATE ITEM
		case "A", "a":
			var number int

			fmt.Print("How many items do you want to add? ")
			fmt.Scan(&number)

			if number <= 0 {
				fmt.Println("Invalid number of items. Please enter a positive number.")
				continue
			}

			for i := 0; i < number; i++ {
				var product Product

				fmt.Printf("Enter item %d name: ", i+1)
				fmt.Scan(&product.Name)

				for {
					fmt.Printf("Enter price %s: ", product.Name)
					fmt.Scan(&product.Price)
					if product.Price < 0 {
						fmt.Println("Invalid price. Please enter a positive number.")
					} else {
						break
					}
				}

				carts = append(carts, product)
			}
			fmt.Println("\nItems added to the cart successfully!")

		// REMOVE ITEM
		case "B", "b":
			if len(carts) == 0 {
				fmt.Println("Cart is empty. No items to remove.")
				continue
			}
			fmt.Print("Enter the index of the item to remove (1 to ", len(carts), "): ")
			for i, product := range carts {
				fmt.Printf("%d. %-15s ₦%.2f\n", i+1, product.Name, product.Price)
			}

			var index int

			fmt.Print("Enter the index of the item to remove (1 to ", len(carts), "): ")
			fmt.Scan(&index)

			index-- // Convert user input to slice index
			if index < 0 || index >= len(carts) {
				fmt.Println("Invalid index. Please try again.")
				continue
			}
			carts = append(carts[:index], carts[index+1:]...)

			fmt.Println("\nItem removed from the cart successfully!")

		// VIEW CART
		case "C", "c":
			if len(carts) == 0 {
				fmt.Println("Cart is empty.")
				continue
			}

			fmt.Println("\n--Your Cart--")
			fmt.Println("\nItems in the cart:")
			for i, product := range carts {
				fmt.Printf("%d. %-15s ₦%.2f\n", i+1, product.Name, product.Price)
			}

		// CALCULATE TOTAL PRICE
		case "D", "d":
			if len(carts) == 0 {
				fmt.Println("Cart is empty. No items to calculate.")
				continue
			}

			var totalPrice float64
			for _, product := range carts {
				totalPrice += product.Price
			}
			fmt.Printf("\nTotal Price: ₦%.2f\n", totalPrice)

		// EXIT
		case "E", "e":
			fmt.Println("Thank you for Shopping with us! Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
