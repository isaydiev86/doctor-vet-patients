package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) GetPreparations(ctx context.Context) ([]*dto.Preparations, error) {
	var preparations []*models.Preparations

	err := pgxscan.Select(ctx, db.DB, &preparations, selectPreparationsSQL)
	if err != nil {
		db.logger.Error("Ошибка получения preparations", err)
		return nil, fmt.Errorf("failed to fetch preparations: %w", err)
	}

	return mapDBPreparationsToDTO(preparations), nil
}

func mapDBPreparationsToDTO(rows []*models.Preparations) []*dto.Preparations {
	preparationsDTO := make([]*dto.Preparations, len(rows))
	for i, row := range rows {
		item := dto.Preparations{
			ID:       row.ID,
			Name:     row.Name.String,
			Dose:     row.Dose.Float64,
			Course:   row.Course.String,
			Category: row.Category.String,
			Option:   row.Option.String,
		}
		preparationsDTO[i] = &item
	}
	return preparationsDTO
}

func (db *DB) GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]*dto.Preparations, error) {
	var preparations []*models.Preparations

	query := `select DISTINCT pr.id, pr.name, pr.dose, pr.course, pr.category, pr.option
				from preparation as pr
				join symptom_relation_preparation as srp on pr.id = srp.preparation_id
				join symptom as s on srp.symptom_id = s.id
			  where s.id = ANY($1);`

	err := pgxscan.Select(ctx, db.DB, &preparations, query, ids)
	if err != nil {
		db.logger.Error("Ошибка получения preparations", err)
		return nil, fmt.Errorf("failed to fetch preparations: %w", err)
	}

	return mapDBPreparationsToDTO(preparations), nil
}
