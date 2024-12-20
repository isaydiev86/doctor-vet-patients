package db

import (
	"context"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) CreatePatient(ctx context.Context, patient dto.Patient) (int64, error) {
	var id int64

	err := db.QueryRow(ctx, createPatientSQL,
		patient.Fio,
		patient.Phone,
		patient.Address,
		patient.Animal,
		patient.Name,
		patient.Breed,
		patient.Gender,
		patient.Age,
		patient.IsNeutered,
	).Scan(&id)

	if err != nil {
		db.logger.Error("Ошибка создания нового пациента", err)
		return 0, fmt.Errorf("failed to create patient: %w", err)
	}

	return id, nil
}

func (db *DB) UpdatePatient(ctx context.Context, patient dto.Patient) error {
	_, err := db.Exec(ctx, updatePatientSQL,
		patient.Fio,
		patient.Phone,
		patient.Address,
		patient.Animal,
		patient.Name,
		patient.Breed,
		patient.Gender,
		patient.Age,
		patient.IsNeutered,
		patient.ID,
	)

	if err != nil {
		db.logger.Error("Ошибка редактирования нового пациента", err)
		return fmt.Errorf("не удалось обновить данные пациента: %w", err)
	}

	return nil
}
