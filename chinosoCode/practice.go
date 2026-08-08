package chinosoCode

import "fmt"

// crud operation for patient

type PatientRecord struct {
	ID   int
	Name string
	Age  int
	Lga  string
}

var patients = []PatientRecord{}

func AddPatient(newPatients []PatientRecord) ([]PatientRecord, []PatientRecord, []PatientRecord) {

	// Make a copy of the current patients
	previousPatients := append([]PatientRecord{}, patients...)

	// Add the new patients
	patients = append(patients, newPatients...)

	return previousPatients, patients, newPatients
}

func ReadPatient() {
	for i := 0; i < len(patients); i++ {
		patient := patients[i]
		fmt.Printf("%d. ID:%d Name:%s Age:%d LGA:%s\n", i+1,
			patient.ID,
			patient.Name,
			patient.Age,
			patient.Lga,
		)
	}
}

func UpdatePatient(index int, newName, newLga string, newAge int) {

	if index < 0 || index >= len(patients) {
		fmt.Println("Invalid patient index.")
		return
	}

	fmt.Println("Update the Student record")

	// Update the Student Record
	patients[index].Name = newName

	patients[index].Age = newAge

	patients[index].Lga = newLga

	fmt.Println("Patient Record updated successfully!")
}

func DeletePatient(index int) {
	if index < 0 || index >= len(patients) {
		fmt.Println("Invalid patient index.")
		return
	}

	// Remove the patient from the slice
	patients = append(patients[:index], patients[index+1:]...)

	fmt.Println("Patient Record deleted successfully!")
}

func PatientCrud() {

	// CREATE patient records
	patientRecord := []PatientRecord{

		{ID: 1, Name: "Ayo", Age: 20, Lga: "Ikeja"},
		{ID: 2, Name: "David", Age: 22, Lga: "Surulere"},
		{ID: 3, Name: "Osiki", Age: 21, Lga: "Yaba"},
	}



	fmt.Println("Adding new patients...")
	previousPatients, patients, newPatients := AddPatient(patientRecord) 

	fmt.Println("Patient Created:", previousPatients, patients, newPatients)

	fmt.Println("Before Adding new patients:")
	ReadPatient()

	// Update the patients slice with the new patients
	UpdatePatient(1, "Ayo Updated", "Ikeja Updated", 21)

	fmt.Println("\nAfter Updating index 1:")
	ReadPatient()

	// DELETE
	// Index 2 means the third patient: Osiki
	DeletePatient(2)

	fmt.Println("\nAfter Deleting index 2:")
	ReadPatient()

}
