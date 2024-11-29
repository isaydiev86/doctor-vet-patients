package db

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (d *DB) CreatePatient(ctx context.Context, patient dto.Patient) error {
	/*
	   	tx, err := db.BeginTxx(ctx, nil)
	   	if err != nil {
	   		return fmt.Errorf("failed to begin transaction: %w", err)
	   	}

	   	defer func() {
	   		if p := recover(); p != nil {
	   			tx.Rollback()
	   			panic(p)
	   		} else if err != nil {
	   			tx.Rollback()
	   		} else {
	   			err = tx.Commit()
	   		}
	   	}()

	   	// Создаем пациента
	   	var patientID int64
	   	patientQuery := `
	           INSERT INTO patient (fio, phone, address, animal, name, breed, gender, is_neutered)
	           VALUES (:fio, :phone, :address, :animal, :name, :breed, :gender, :is_neutered)
	           RETURNING id
	       `
	   	row := tx.NamedQueryContext(ctx, patientQuery, patient)
	   	if row.Next() {
	   		err = row.Scan(&patientID)
	   		if err != nil {
	   			return fmt.Errorf("failed to insert patient: %w", err)
	   		}
	   	} else {
	   		return fmt.Errorf("failed to fetch created patient ID")
	   	}

	   	// Создаем запись о лечении
	   	treatmentQuery := `
	           INSERT INTO treatments (patient_id, status, created_at, updated_at, is_active)
	           VALUES ($1, $2, NOW(), NOW(), 1)
	       `
	   	_, err = tx.ExecContext(ctx, treatmentQuery, patientID, models.InLine.String())
	   	if err != nil {
	   		return fmt.Errorf("failed to insert treatment: %w", err)
	   	}
	   /**/
	return nil
}
