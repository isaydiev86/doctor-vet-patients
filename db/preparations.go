package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (db *DB) CreatePreparations(ctx context.Context, pr dto.PreparationsAdd) error {
	query := `
		INSERT INTO preparation (name, dose, course, category, option)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (name) DO UPDATE
		SET dose = EXCLUDED.dose,
		    course = EXCLUDED.course,
		    category = EXCLUDED.category,
		    option = EXCLUDED.option;
	`
	_, err := db.Exec(ctx, query, pr.Name, pr.Dose, pr.Course, pr.Category, pr.Option)

	if err != nil {
		db.logger.Error("failed to create preparation", err)
		return fmt.Errorf("failed to create preparation: %w", err)
	}

	return nil
}

func (db *DB) GetPreparations(ctx context.Context) ([]dto.Preparations, error) {
	var preparations []*models.Preparations

	err := pgxscan.Select(ctx, db.DB, &preparations, selectPreparationsSQL)
	if err != nil {
		db.logger.Error("Ошибка получения preparations", err)
		return nil, fmt.Errorf("failed to fetch preparations: %w", err)
	}

	return mapDBPreparationsToDTO(preparations), nil
}

func mapDBPreparationsToDTO(rows []*models.Preparations) []dto.Preparations {
	preparationsDTO := make([]dto.Preparations, len(rows))
	for i, row := range rows {
		item := dto.Preparations{
			ID:       row.ID,
			Name:     row.Name.String,
			Dose:     row.Dose.Float64,
			Course:   row.Course.String,
			Category: row.Category.String,
			Option:   row.Option.String,
		}
		preparationsDTO[i] = item
	}
	return preparationsDTO
}

func (db *DB) GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]dto.Preparations, error) {
	var preparations []*models.Preparations

	// группируем по популярности
	query := `
		WITH ranked_preparations AS (
			SELECT 
				pr.id, 
				pr.name, 
				pr.dose, 
				pr.course, 
				pr.category, 
				pr.option,
				ROW_NUMBER() OVER (PARTITION BY pr.category ORDER BY pr.popularity DESC) as rank
			FROM preparation as pr
			JOIN symptom_relation_preparation as srp ON pr.id = srp.preparation_id
			JOIN symptom as s ON srp.symptom_id = s.id
			WHERE s.id = ANY($1)
		)
		SELECT id, name, dose, course, category, option
		FROM ranked_preparations
		WHERE rank = 1;
	`

	err := pgxscan.Select(ctx, db.DB, &preparations, query, ids)
	if err != nil {
		db.logger.Error("Ошибка получения preparations", err)
		return nil, fmt.Errorf("failed to fetch preparations: %w", err)
	}

	return mapDBPreparationsToDTO(preparations), nil
}
