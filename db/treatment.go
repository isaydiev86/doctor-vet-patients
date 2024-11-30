package db

import (
	"context"
	"encoding/json"
	"fmt"

	"doctor-vet-patients/db/models"
	"doctor-vet-patients/internal/dto"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (db *DB) CreateTreatment(ctx context.Context, patientID int64) (int64, error) {
	treatmentQuery := `
		INSERT INTO treatment (patient_id, status, created_at, updated_at, is_active)
		VALUES ($1, $2, NOW(), NOW(), 1)
		RETURNING id
	`
	var treatmentID int64
	err := db.QueryRow(ctx, treatmentQuery, patientID, models.InLine.String()).Scan(&treatmentID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert treatment: %w", err)
	}

	return treatmentID, nil
}

func (db *DB) GetTreatments(ctx context.Context) ([]*dto.Treatment, error) {
	var treatments []*models.TreatmentRow

	err := pgxscan.Select(ctx, db.DB, &treatments, selectTreatmentsSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch treatments: %w", err)
	}

	return mapTreatmentDBtoDTO(treatments), nil
}

func mapTreatmentDBtoDTO(rows []*models.TreatmentRow) []*dto.Treatment {
	treatments := make([]*dto.Treatment, len(rows))

	for i, row := range rows {
		item := dto.Treatment{
			ID:          row.ID,
			PatientID:   row.PatientID,
			DoctorID:    row.DoctorID.String,
			Temperature: row.Temperature.Float64,
			Status:      row.Status.String,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			BeginAt:     row.BeginAt.Time,
			EndAt:       row.EndAt.Time,
			Comment:     row.Comment.String,
			IsActive:    row.IsActive,
			Weight:      row.Weight.Float64,
			Patient: dto.Patient{
				ID:         row.Patient.ID,
				Fio:        row.Patient.Fio.String,
				Phone:      row.Patient.Fio.String,
				Address:    row.Patient.Address.String,
				Animal:     row.Patient.Animal.String,
				Name:       row.Patient.Name.String,
				Breed:      row.Patient.Breed.String,
				Age:        row.Patient.Age.Float64,
				Gender:     row.Patient.Gender.String,
				IsNeutered: row.Patient.IsNeutered,
			},
		}
		treatments[i] = &item
	}

	return treatments
}

func (db *DB) GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error) {
	var treatment models.TreatmentDetailRow

	err := pgxscan.Get(ctx, db.DB, &treatment, selectTreatmentDetailSQL, id)
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, fmt.Errorf("лечение с id %d не найдено", id)
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

func mapTreatmentDetailDBtoDTO(row *models.TreatmentDetailRow) *dto.TreatmentDetail {

	prescription := make([]dto.Prescription, len(row.Prescription))

	for i, v := range row.Prescription {
		item := dto.Prescription{
			ID:          v.ID,
			TreatmentID: v.TreatmentID,
			Preparation: v.Preparation.String,
			Dose:        v.Dose.String,
			Course:      v.Course.String,
			Amount:      v.Amount.String,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		prescription[i] = item
	}

	treatmentDetail := &dto.TreatmentDetail{
		ID:          row.ID,
		PatientID:   row.PatientID,
		DoctorID:    row.DoctorID.String,
		Temperature: row.Temperature.Float64,
		Status:      row.Status.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		BeginAt:     row.BeginAt.Time,
		EndAt:       row.EndAt.Time,
		Comment:     row.Comment.String,
		IsActive:    row.IsActive,
		Weight:      row.Weight.Float64,
		Patient: dto.Patient{
			ID:         row.Patient.ID,
			Fio:        row.Patient.Fio.String,
			Phone:      row.Patient.Fio.String,
			Address:    row.Patient.Address.String,
			Animal:     row.Patient.Animal.String,
			Name:       row.Patient.Name.String,
			Breed:      row.Patient.Breed.String,
			Age:        row.Patient.Age.Float64,
			Gender:     row.Patient.Gender.String,
			IsNeutered: row.Patient.IsNeutered,
		},
		Prescription: prescription,
	}

	return treatmentDetail
}
