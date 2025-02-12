package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/internal/errors"
)

func (db *DB) GetTreatmentForUser(ctx context.Context, userId string) (*dto.TreatmentDetail, error) {
	var treatment models.TreatmentDetailRow

	query := `
		SELECT 
            t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at,
            t.comment, t.is_active, t.weight, t.add_info,
            p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone",
            p.animal AS "patient.animal", p.name AS "patient.name", p.breed AS "patient.breed", p.gender AS "patient.gender",
            p.age AS "patient.age", p.is_neutered AS "patient.is_neutered",
            COALESCE(
                json_agg(
                    json_build_object(
                        'id', pr.id,
                        'treatment_id', pr.treatment_id,
                        'name', pr.name,
                        'course', pr.course,
                        'dose', pr.dose,
                        'category', pr.category,
                        'option', pr.option,
                        'created_at', pr.created_at,
                        'updated_at', pr.updated_at
                    )
                ) FILTER (WHERE pr.id IS NOT NULL), '[]'
            ) AS prescriptions
        FROM 
            treatment t
        LEFT JOIN 
            patient p ON t.patient_id = p.id
        LEFT JOIN 
            prescription pr ON t.id = pr.treatment_id
        WHERE 
            t.doctor_id = $1 and t.status = 'процесс'
		GROUP BY 
            t.id, p.id LIMIT 1;
`

	err := pgxscan.Get(ctx, db.DB, &treatment, query, userId)
	if err != nil {
		if pgxscan.NotFound(err) {
			db.logger.Error("лечение с id %s не найдено", err)
			return nil, errors.ErrNotFound
		}
		db.logger.Error("не удалось получить детали лечения", err)
		return nil, fmt.Errorf("не удалось получить детали лечения: %w", err)
	}

	// Парсим поле prescriptions из JSON в слайс структур Prescription
	err = json.Unmarshal(treatment.PrescriptionsJSON, &treatment.Prescription)
	if err != nil {
		db.logger.Error("не удалось распарсить рецепты", err)
		return nil, fmt.Errorf("не удалось распарсить рецепты: %w", err)
	}

	return mapTreatmentDetailDBtoDTO(&treatment), nil
}
