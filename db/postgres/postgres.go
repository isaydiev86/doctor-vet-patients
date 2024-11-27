package postgres

import (
	"context"

	"doctor-vet-patients/db"
	"doctor-vet-patients/db/models"
	"doctor-vet-patients/internal/dto"
)

type repo struct {
}

func NewRepoPostgres() db.IStorage {
	return &repo{}
}

func (r *repo) GetPatients(ctx context.Context) ([]dto.Patient, error) {
	return getPatients(), nil
}

func getPatients() []dto.Patient {
	patients := make([]dto.Patient, 0)
	patients = append(patients, dto.Patient{
		ID:          1,
		DoctorID:    "UUID1",
		Fio:         "Test1",
		Phone:       "1111",
		Address:     "Address1",
		Animal:      "dog",
		Name:        "black",
		Breed:       "breed1",
		Age:         2,
		Weight:      4.5,
		Temperature: 37,
		Gender:      "мужской",
		Status:      models.InProcess.String(),
		IsNeutered:  false,
	})
	patients = append(patients, dto.Patient{
		ID:          2,
		DoctorID:    "UUID2",
		Fio:         "Test2222",
		Phone:       "33333",
		Address:     "Address155555",
		Animal:      "cat",
		Name:        "Мурзик",
		Breed:       "breed555",
		Age:         3,
		Weight:      1.5,
		Temperature: 39,
		Gender:      "женский",
		Status:      models.InProcess.String(),
		IsNeutered:  true,
	})

	return patients
}
