package db

import (
	"context"
	"fmt"

	"doctor-vet-patients/db/models"
	"doctor-vet-patients/internal/dto"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

func (db *DB) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	var references []*models.Reference

	err := pgxscan.Select(ctx, db.DB, &references, selectReferenceSQL, typeQuery)
	if err != nil {
		db.logger.Error("Ошибка получения справочника", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch references: %w", err)
	}

	return mapDBReferenceToDTO(references), nil
}

func mapDBReferenceToDTO(rows []*models.Reference) []*dto.Reference {
	referenceDTO := make([]*dto.Reference, len(rows))
	for i, row := range rows {
		item := dto.Reference{
			ID:   row.ID,
			Name: row.Name.String,
			Type: row.Type.String,
		}
		referenceDTO[i] = &item
	}
	return referenceDTO
}
