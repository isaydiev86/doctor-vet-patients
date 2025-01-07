package db

import (
	"context"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) UpdateStatusTreatment(ctx context.Context, treatment dto.TreatmentUpdateStatus) error {
	treatmentQuery := `
		UPDATE treatment 
		SET status = $1, updated_at = now()
    	WHERE id = $2;`

	_, err := db.Exec(ctx, treatmentQuery, treatment.Status, treatment.ID)
	if err != nil {
		db.logger.Error("failed to update status treatment", err)
		return fmt.Errorf("failed to update status treatment: %w", err)
	}
	return nil
}
