package postgres

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (d *DB) GetTreatments(ctx context.Context) ([]*dto.Treatment, error) {
	return nil, nil
}
