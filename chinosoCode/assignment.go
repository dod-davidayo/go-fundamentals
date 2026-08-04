package chinosoCode

import "fmt"

type StudentRecord struct {
	ID   int
	Name string
	Age  int
	Lga  string
}

var students = []StudentRecord{}

func AddStudent(newStudents []StudentRecord) ([]StudentRecord, []StudentRecord, []StudentRecord) {

	previousStudents := append([]StudentRecord{}, students...)

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
		fmt.Printf("%d. ID:%d Name:%s Age:%d LGA:%s\n",
			i+1,
			student.ID,
			student.Name,
			student.Age,
			student.Lga,
		)

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

	students[index-1].Name = newName

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

	//students = append(students, newStudents...)

	students = append(students[:index-1], students[index:]...)
	fmt.Println("Student deleted successfully!")
}

func StudentManagement() {
	var number int

	fmt.Print("How many students do you want to add? ")
	fmt.Scan(&number)

	newStudents := []StudentRecord{}

	for i := 0; i < number; i++ {
		var student StudentRecord

		fmt.Print("Enter ID: ")
		fmt.Scan(&student.ID)

		fmt.Print("Enter Name: ")
		fmt.Scan(&student.Name)

		fmt.Print("Enter Age: ")
		fmt.Scan(&student.Age)

		fmt.Print("Enter LGA: ")
		fmt.Scan(&student.Lga)

		newStudents = append(newStudents, student)
	}

	oldStudents, updatedStudents, addedStudents := AddStudent(newStudents)

	fmt.Println("Old Students:", oldStudents)
	fmt.Println("Updated Students:", updatedStudents)
	fmt.Println("New Students:", addedStudents)
}
