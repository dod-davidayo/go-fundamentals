package interfaces

import "fmt"

// Payment interfaces

//Any type that has a Pay() method
// Automatically satisfies the Payment interface.

type Payment interface{
	Pay()
}

// Cash Payment
type Cash struct {
	Amount float64
}

// Cash Implement the Payment Interface
func (cash Cash) Pay() {
	fmt.Printf("Paid #%.2f with Cash\n", cash.Amount)
}

// Card Payment
type Card struct{
	Amount float64
}

// Card implement the  Payment interface
func (card Card) Pay() {
	fmt.Printf("Paid #%.2f with card\n", card.Amount)
}

// Transfer payment
type Transfer struct {
	Amount float64
}

// Transfer implements the Payment interface
func (transfer Transfer) Pay(){
	fmt.Printf("Paid #%.2f through Bank Transfer\n", transfer.Amount)
}

func PaymentDemo() {
	// Create instances of different payment methods
	cashPayment := Cash{Amount: 100.0}
	cardPayment := Card{Amount: 200.0}
	transferPayment := Transfer{Amount: 300.0}

	// Create a slice of Payment interface
	payments := []Payment{cashPayment, cardPayment, transferPayment}

	// call for each payment method
	for _, payment := range payments {
		payment.Pay()
	}

}
