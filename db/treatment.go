package db

import (
	"context"
	"fmt"

	"doctor-vet-patients/internal/dto"
)

func (d *DB) GetTreatments(ctx context.Context) ([]*dto.Treatment, error) {
	var treatments []*dto.Treatment

	query := `
        SELECT 
            t.id AS id,
            t.patient_id AS patient_id,
            t.doctor_id AS doctor_id,
            t.temperature AS temperature,
            t.status AS status,
            t.created_at AS created_at,
            t.updated_at AS updated_at,
            t.begin_at AS begin_at,
            t.end_at AS end_at,
            t.comment AS comment,
            t.is_active AS is_active,
            t.age AS age,
            t.weight AS weight,
            p.id AS "patient.id",
            p.fio AS "patient.fio",
            p.phone AS "patient.phone",
            p.address AS "patient.address",
            p.animal AS "patient.animal",
            p.name AS "patient.name",
            p.breed AS "patient.breed",
            p.gender AS "patient.gender",
            p.is_neutered AS "patient.is_neutered"
        FROM 
            treatments t
        LEFT JOIN 
            patients p 
        ON 
            t.patient_id = p.id;
    `

	// Выполняем запрос с sqlx
	err := d.DB.SelectContext(ctx, &treatments, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch treatments: %w", err)
	}

	return treatments, nil
}

func (d *DB) GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error) {
	var treatment dto.TreatmentDetail
	var prescriptions []dto.Prescription

	query := `
        SELECT 
            t.id AS id,
            t.patient_id AS patient_id,
            t.doctor_id AS doctor_id,
            t.temperature AS temperature,
            t.status AS status,
            t.created_at AS created_at,
            t.updated_at AS updated_at,
            t.begin_at AS begin_at,
            t.end_at AS end_at,
            t.comment AS comment,
            t.is_active AS is_active,
            t.age AS age,
            t.weight AS weight,
            p.id AS "patient.id",
            p.fio AS "patient.fio",
            p.phone AS "patient.phone",
            p.address AS "patient.address",
            p.animal AS "patient.animal",
            p.name AS "patient.name",
            p.breed AS "patient.breed",
            p.gender AS "patient.gender",
            p.is_neutered AS "patient.is_neutered"
        FROM 
            treatments t
        LEFT JOIN 
            patients p 
        ON 
            t.patient_id = p.id
        WHERE 
            t.id = $1;
    `

	err := d.DB.GetContext(ctx, &treatment, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch treatment details: %w", err)
	}

	// Получаем рецепты, связанные с лечением
	prescriptionQuery := `
        SELECT 
            id,
            treatment_id,
            preparation,
            course,
            dose,
            amount,
            created_at
        FROM 
            prescriptions
        WHERE 
            treatment_id = $1;
    `

	err = d.DB.SelectContext(ctx, &prescriptions, prescriptionQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prescriptions: %w", err)
	}

	treatment.Prescription = prescriptions

	return &treatment, nil
}
