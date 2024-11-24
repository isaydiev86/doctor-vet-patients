package mapper

import (
	"doctor-vet-patients/internal/dto"
	"doctor-vet-patients/transport/models"
)

func GetDtoPatientOfApi(patient models.Patient) dto.Patient {
	return dto.Patient{
		DoctorID:    patient.DoctorID,
		Fio:         patient.Fio,
		Phone:       patient.Phone,
		Address:     patient.Address,
		Animal:      patient.Animal,
		Name:        patient.Name,
		Breed:       patient.Breed,
		Age:         patient.Age,
		Weight:      patient.Weight,
		Temperature: patient.Temperature,
		Gender:      patient.Gender,
		Status:      patient.Status,
		IsNeutered:  patient.IsNeutered,
	}
}
