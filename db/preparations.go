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

func (db *DB) GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]*dto.PreparationsWithSimilar, error) {
	var preparations []*models.PreparationsToSymptoms

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
				pr.popularity,
				ROW_NUMBER() OVER (PARTITION BY pr.category ORDER BY pr.popularity DESC) as rank
			FROM preparation as pr
			JOIN symptom_relation_preparation as srp ON pr.id = srp.preparation_id
			JOIN symptom as s ON srp.symptom_id = s.id
			WHERE s.id = ANY($1)
		),
		similar_preparations AS (
			SELECT
				rp.category,
				jsonb_agg(DISTINCT jsonb_build_object(
					'id', rp.id,
					'name', rp.name
				)) AS similar
			FROM ranked_preparations rp
			WHERE rp.rank > 1
			GROUP BY rp.category
		)
		SELECT 
			rp.id, 
			rp.name, 
			rp.dose, 
			rp.course, 
			rp.category, 
			rp.option,
			COALESCE(sp.similar, '[]'::jsonb) AS similar
		FROM ranked_preparations rp
		LEFT JOIN similar_preparations sp ON rp.category = sp.category
		WHERE rp.rank = 1;
	`

	err := pgxscan.Select(ctx, db.DB, &preparations, query, ids)
	if err != nil {
		db.logger.Error("Ошибка получения preparations", err)
		return nil, fmt.Errorf("failed to fetch preparations: %w", err)
	}

	return mapDBPreparationsWithSimilarToDTO(preparations), nil
}

func mapDBPreparationsWithSimilarToDTO(rows []*models.PreparationsToSymptoms) []*dto.PreparationsWithSimilar {
	preparationsDTO := make([]*dto.PreparationsWithSimilar, len(rows))
	for i, row := range rows {
		item := dto.PreparationsWithSimilar{
			ID:       row.ID,
			Name:     row.Name.String,
			Dose:     row.Dose.Float64,
			Course:   row.Course.String,
			Category: row.Category.String,
			Option:   row.Option.String,
			Similar:  mapSimilarDBToDto(row.Similar),
		}
		preparationsDTO[i] = &item
	}
	return preparationsDTO
}

func mapSimilarDBToDto(rows []models.NameRow) []dto.NameResponse {
	similarDTO := make([]dto.NameResponse, len(rows))
	for i, row := range rows {
		item := dto.NameResponse{
			ID:   row.ID,
			Name: row.Name,
		}
		similarDTO[i] = item
	}
	return similarDTO
}
