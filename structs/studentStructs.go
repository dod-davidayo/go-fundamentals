package structs


//packa

import "fmt"

// Question 1

// Create a Student struct.

// Fields:

// ID
// Name
// Age
// Department

// Create five students and display them.


type Student struct {
	ID     int
	Name   string
	Age    string
	Dept   string

}

func CrudStudent() {

	// create 5 student and display 
	var books [5]Student{
		{ID: 1, Name: "Ayo", Age: 20, Dept: "Computer Science"},
		{ID: 2, Name: "David", Age: 22, Dept: "Mathematics"},
		{ID: 3, Name: "Osiki", Age: 21, Dept: "Biolgy"},
		{ID: 4, Name: "Patience", Age: 21, Dept: "Biolgy"},
		{ID: 5, Name: "Faith", Age: 21, Dept: "Biolgy"},	
	}

	// Display all students
	fmt.Println("Student Records:")

	for i, student := range students {
		fmt.Printf("%d. ID: %d | Name: %s | Age: %d | Department: %s\n",
		    i+1,
			student.ID,
			student.Name,
			student.Age,
			student.Dept,
	)
	}

}