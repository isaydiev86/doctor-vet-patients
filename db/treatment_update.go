package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

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

func (db *DB) GetTreatmentForUser(ctx context.Context, userId string) (*dto.TreatmentDetail, error) {
	var treatment models.TreatmentDetailRow

	query := `
		SELECT 
            t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at,
            t.comment, t.is_active, t.weight,
            p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone", p.address AS "patient.address",
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
            t.doctor_id = $1 and t.status = 'в процессе'
		GROUP BY 
            t.id, p.id;
`

	err := pgxscan.Get(ctx, db.DB, &treatment, query, userId)
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, fmt.Errorf("лечение с id %s не найдено", userId)
		}
		return nil, fmt.Errorf("не удалось получить детали лечения: %w", err)
	}

	// Парсим поле prescriptions из JSON в слайс структур Prescription
	err = json.Unmarshal(treatment.PrescriptionsJSON, &treatment.Prescription)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить рецепты: %w", err)
	}

	return mapTreatmentDetailDBtoDTO(&treatment), nil
}
