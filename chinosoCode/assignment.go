package chinosoCode

import "fmt"

type StudentRecord struct {
	ID   int
	Name string
	Age  int
	Lga  string
}

var students = []StudentRecord{}

// ADD STUDENT
func AddStudent(newStudents []StudentRecord) ([]StudentRecord, []StudentRecord, []StudentRecord) {

	// Make a copy of the current students
	previousStudents := append([]StudentRecord{}, students...)

	// Add the new students
	students = append(students, newStudents...)

	return previousStudents, students, newStudents
}

// READ STUDENTS
func ReadStudent() {

	if len(students) == 0 {
		fmt.Println("No Student Found.")
		return
	}

	fmt.Println("\n--- Student List ---")

	for i, student := range students {
		fmt.Printf(
			"%d. ID:%d Name:%s Age:%d LGA:%s\n",
			i+1,
			student.ID,
			student.Name,
			student.Age,
			student.Lga,
		)
	}
}

// UPDATE STUDENT
func UpdatedStudent(index int, newName string) {

	if index < 0 || index >= len(students) {
		fmt.Println("Invalid student index.")
		return
	}

	// Update the student's name
	students[index].Name = newName

	fmt.Println("Student updated successfully!")
}

// DELETE STUDENT
func DeletedStudent(index int) {

	if index < 0 || index >= len(students) {
		fmt.Println("Invalid student index.")
		return
	}

	// Remove the student using slicing
	students = append(students[:index], students[index+1:]...)

	fmt.Println("Student deleted successfully!")
}

// STUDENT MANAGEMENT
func StudentManagement() {

	// Create students directly instead of using fmt.Scan()
	newStudents := []StudentRecord{
		{
			ID:   1,
			Name: "David",
			Age:  22,
			Lga:  "Ethiope East",
		},
		{
			ID:   2,
			Name: "Ayo",
			Age:  21,
			Lga:  "Oshimili South",
		},
		{
			ID:   3,
			Name: "Osiki",
			Age:  23,
			Lga:  "Ethiope East",
		},
	}

	// ADD
	oldStudents, updatedStudents, addedStudents := AddStudent(newStudents)

	fmt.Println("Old Students:", oldStudents)
	fmt.Println("Updated Students:", updatedStudents)
	fmt.Println("New Students:", addedStudents)

	// READ
	ReadStudent()

	// UPDATE
	// Index 1 means the second student: Ayo
	UpdatedStudent(1, "Ayo Updated")

	fmt.Println("\nAfter Updating index 1:")
	ReadStudent()

	// DELETE
	// Index 2 means the third student: Osiki
	DeletedStudent(2)

	fmt.Println("\nAfter Deleting index 2:")
	ReadStudent()
}
