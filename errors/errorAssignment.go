package errors

// Create a login function.

// Return an error if:

// username is empty
// password is empty



import (
	"errors"
	"fmt"
)

// login checks  if username and password are provided.
func login(username, password string) error {

	// check if username is empty
	if username == "" {
		return errors.New("username cannot be empty")
	}

	// check if password is empty
	if password == "" {
		return errors.New("password cannot be empty")
	}

	// Login is successful
	return nil

}

func LoginDemo() {
	// Test cases
	err := login("", "password123")

	if err != nil {
		fmt.Println("Error:", err)

	} else {

		fmt.Println("Login successful")

	}

	
}