package chinosocode

import "fmt"

var students = []string{}

func AddStudent(newStudents []string) ([]string, []string, []string) {

	previousStudents := append([]string{}, students...)

	students = append(students, newStudents...)

	return previousStudents, students, newStudents
}

func ReadStudent() {
	if len(students) == 0 {
		fmt.Println("No Student Found.")
		return
	}

	fmt.Println("\n --- Student List ------")

	for i, student := range students {
		fmt.Printf("%d. %s\n", i+1, student)

	}
}

func UpdatedStudent() {

	var index int
	var newName string

	fmt.Print("Enter student number to update: ")
	fmt.Scan(&index)

	if index < 1 || index > len(students) {
		fmt.Println("Invalid student number.")
		return
	}

	fmt.Print("Enter new student name: ")
	fmt.Scan(&newName)

	students[index-1] = newName

	fmt.Println("Student updated successfully!")
}

func DeletedStudent() {
	var index int

	fmt.Print("Enter student number to delete: ")
	fmt.Scan(&index)

	if index < 1 || index > len(students) {
		fmt.Println("Invalid student number.")
		return
	}

	students = append(students[:index-1], students[index:]...)
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
