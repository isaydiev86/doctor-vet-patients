package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) UpdateTreatment(ctx context.Context, treatment dto.TreatmentUpdateToUser) error {
	addInfoJSON, err := json.Marshal(treatment.AddInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal add_info: %w", err)
	}

	treatmentQuery := `
		UPDATE treatment 
		SET doctor_id = $1, temperature = $2, weight = $3, comment = $4, status = $5, add_info =$6, updated_at = now()
    	WHERE id = $7;`

	_, err = db.Exec(ctx, treatmentQuery, treatment.DoctorID, treatment.Temperature, treatment.Weight, treatment.Comment,
		models.InPayment.String(), addInfoJSON, treatment.ID)
	if err != nil {
		return fmt.Errorf("failed to update treatment: %w", err)
	}
	return nil
}

func (db *DB) UpdateTreatmentForUser(ctx context.Context, treatment dto.TreatmentSendForUser) error {
	treatmentQuery := `
		UPDATE treatment 
		SET doctor_id = $1, status = $2, updated_at = now(), begin_at = now()
    	WHERE id = $3;`

	_, err := db.Exec(ctx, treatmentQuery, treatment.DoctorID, models.InProcess.String(), treatment.ID)
	if err != nil {
		return fmt.Errorf("failed to update treatment: %w", err)
	}
	return nil
}
