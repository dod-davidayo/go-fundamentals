package chinosocode

import "fmt"

var students = []string{}

func AddStudent(newStudents []string) ([]string, []string, []string) {

	previousStudents := append([]string{}, students...)

	students = append(students, newStudents...)

	return previousStudents, students, newStudents
}

func StudentManagement() {
	var number int

	fmt.Print("How many students do you want to add? ")
	fmt.Scan(&number)

	newStudents := []string{}

	for i := 0; i < number; i++ {
		var name string

		fmt.Printf("Enter student %d name: ", i+1)
		fmt.Scan(&name)

		newStudents = append(newStudents, name)
	}

	oldStudents, updatedStudents, addedStudents := AddStudent(newStudents)

	fmt.Println("Old Students:", oldStudents)
	fmt.Println("Updated Students:", updatedStudents)
	fmt.Println("New Students:", addedStudents)
}
