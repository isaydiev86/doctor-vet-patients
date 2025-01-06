package db

import (
	"context"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/jackc/pgx/v5"
)

func (db *DB) AddPrescriptionsToTreatment(ctx context.Context, treatmentID int64, list []dto.PrescriptionForUpdate) error {
	prescriptions := mapPrescriptionsDTOtoSQL(treatmentID, list)

	batch := &pgx.Batch{}

	query := `
		INSERT INTO prescription (treatment_id, name, dose, course, category, option, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, now(), now())
	`

	// Добавляем каждую запись в batch
	for _, prescription := range prescriptions {
		batch.Queue(query, prescription.TreatmentID, prescription.Name, prescription.Dose, prescription.Course, prescription.Category, prescription.Option)
	}

	br := db.SendBatch(ctx, batch)
	defer func(br pgx.BatchResults) {
		err := br.Close()
		if err != nil {
			db.logger.Error("failed close batch", err)
		}
	}(br)

	if err := br.Close(); err != nil {
		return fmt.Errorf("unable to execute batch insert: %w", err)
	}

	return nil
}

func mapPrescriptionsDTOtoSQL(treatmentID int64, dto []dto.PrescriptionForUpdate) []models.PrescriptionRow {
	rows := make([]models.PrescriptionRow, len(dto))

	for i, p := range dto {
		rows[i] = models.PrescriptionRow{
			TreatmentID: treatmentID,
			Name:        utils.ValidNullString(p.Name),
			Dose:        utils.ValidNullFloat64(p.Dose),
			Course:      utils.ValidNullString(p.Course),
			Category:    utils.ValidNullString(p.Category),
			Option:      utils.ValidNullString(p.Option),
		}
	}

	return rows
}
