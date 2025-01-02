package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) GetSymptoms(ctx context.Context) ([]*dto.Symptoms, error) {
	var symptoms []*models.Symptoms

	err := pgxscan.Select(ctx, db.DB, &symptoms, selectSymptomsSQL)
	if err != nil {
		db.logger.Error("Ошибка получения справочника", err)
		return nil, fmt.Errorf("failed to fetch references: %w", err)
	}

	return mapDBSymptomsToDTO(symptoms), nil
}

func mapDBSymptomsToDTO(rows []*models.Symptoms) []*dto.Symptoms {
	symptomsDTO := make([]*dto.Symptoms, len(rows))
	for i, row := range rows {
		item := dto.Symptoms{
			ID:   row.ID,
			Name: row.Name.String,
		}
		symptomsDTO[i] = &item
	}
	return symptomsDTO
}