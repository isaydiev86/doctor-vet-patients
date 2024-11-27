package postgres

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (d *DB) CreatePatient(ctx context.Context, patient dto.Patient) error {
	query := `
		INSERT INTO patient (fio, phone, address, animal, name, breed, gender, is_neutered)
		VALUES (:fio, :phone, :address, :animal, :name, :breed, :gender, :is_neutered)
		RETURNING id
	`

	_, err := d.db.NamedExecContext(ctx, query, patient)
	if err != nil {
		return err
	}

	return nil
}
