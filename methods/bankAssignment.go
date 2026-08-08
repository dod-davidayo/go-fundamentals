package methods

import "fmt"

// Question 2

// Create a BankAccount struct.

// Methods:

// Deposit()

// Withdraw()

// DisplayBalance()

type BankAccount struct {
	AccountName    string
	AccountNumber  string
	Balance        float64 
}

// Deposit adds money to the account.
func (account *BankAccount) Deposit(amount float64) {

	if amount <= 0 {
		fmt.Println("Deposit amount must be greater than zero.")
		return
	}

	account.Balance += amount

	fmt.Printf("Successfully deposited #%.2f\n", amount)
}

// withdraw and remove money from account.
func (account *BankAccount) Withdraw(amount float64) {

	if amount <= 0 {
		fmt.Println("Withdrawal amount must be greater than zero.")
		return
	}

	if amount > account.Balance {
		fmt.Println("Insufficient balance.")
		return
	}

	account.Balance -= amount

	fmt.Printf("Successfully withdrew #%.2f\n", amount)
}

// DisplayBalance displays the current account balance.
func (account *BankAccount) DisplayBalance() {

	fmt.Printf("\nAccount Name: %s\n", account.AccountName)
	fmt.Printf("\nAccount Number: %s\n", account.AccountNumber)
	fmt.Printf("\nCurrent Balance: %s\n", account.Balance)
}

func BankAccountDemo() {

	// Create a bank account.
	account := BankAccount{
		AccountName: "David",
		AccountNumber: "1234567890",
		Balance: 10000,
	}

	// Display inital balance
	account.DisplayBalance()

	// Deposit balance after deposit.
	account.Deposit(5000)

	// Display balance after deposit
	account.DisplayBalance()

	// Withdraw Money.
	account.Withdraw(3000)

	// Display Final balance
	account.DisplayBalance()
}