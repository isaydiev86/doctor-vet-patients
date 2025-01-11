package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/db/models"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
)

func (db *DB) CreateTreatment(ctx context.Context, patientID int64) (int64, error) {
	treatmentQuery := `
		INSERT INTO treatment (patient_id, status, created_at, updated_at, is_active)
		VALUES ($1, $2, NOW(), NOW(), 1)
		RETURNING id;
	`
	var treatmentID int64
	err := db.QueryRow(ctx, treatmentQuery, patientID, models.Wait.String()).Scan(&treatmentID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert treatment: %w", err)
	}

	return treatmentID, nil
}

func (db *DB) GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error) {
	var treatments []*models.TreatmentRow

	query := `
	SELECT
		t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at,
		t.comment, t.is_active, t.weight,
		p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone",
		p.age AS "patient.age", p.animal AS "patient.animal", p.name AS "patient.name",
		p.breed AS "patient.breed", p.gender AS "patient.gender", p.is_neutered AS "patient.is_neutered"
	FROM treatment t
	LEFT JOIN
		 patient p ON t.patient_id = p.id
	WHERE
		($1::TEXT IS NULL OR p.fio ILIKE '%' || $1 || '%') AND
		($2::TEXT IS NULL OR p.name ILIKE '%' || $2 || '%') AND
		($3::TEXT IS NULL OR t.status = $3) AND 
		($4::DATE IS NULL OR DATE(t.created_at) = $4::DATE)
	ORDER BY
		t.created_at DESC
	LIMIT $5 OFFSET $6;`

	err := pgxscan.Select(ctx, db.DB, &treatments, query,
		utils.NilIfEmpty(filter.Fio),
		utils.NilIfEmpty(filter.Name),
		utils.NilIfEmpty(filter.Status),
		utils.NilIfEmpty(filter.Date),
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		db.logger.Error("db on GetTreatments", err)
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
			BeginAt:     utils.ToTimePtr(row.BeginAt),
			EndAt:       utils.ToTimePtr(row.EndAt),
			Comment:     row.Comment.String,
			IsActive:    row.IsActive,
			Weight:      row.Weight.Float64,
			Patient: dto.Patient{
				ID:         row.Patient.ID,
				Fio:        row.Patient.Fio.String,
				Phone:      row.Patient.Phone.String,
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

	// Парсим поле add_info из JSON в слайс структур AddInfo
	err := json.Unmarshal(row.AddInfoJSON, &row.AddInfo)
	if err != nil {
		return nil
	}

	for i, v := range row.Prescription {
		item := dto.Prescription{
			ID:          v.ID,
			TreatmentID: v.TreatmentID,
			Name:        v.Name.String,
			Dose:        v.Dose.Float64,
			Course:      v.Course.String,
			Category:    v.Category.String,
			Option:      v.Option.String,
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
		BeginAt:     utils.ToTimePtr(row.BeginAt),
		EndAt:       utils.ToTimePtr(row.EndAt),
		Comment:     row.Comment.String,
		IsActive:    row.IsActive,
		Weight:      row.Weight.Float64,
		Patient: dto.Patient{
			ID:         row.Patient.ID,
			Fio:        row.Patient.Fio.String,
			Phone:      row.Patient.Phone.String,
			Animal:     row.Patient.Animal.String,
			Name:       row.Patient.Name.String,
			Breed:      row.Patient.Breed.String,
			Age:        row.Patient.Age.Float64,
			Gender:     row.Patient.Gender.String,
			IsNeutered: row.Patient.IsNeutered,
		},
		Prescription: prescription,
		AddInfo:      mapAddInfoToDTO(row.AddInfo),
	}

	return treatmentDetail
}

func mapAddInfoToDTO(row []models.AddInfo) []dto.AddInfo {
	addInfo := make([]dto.AddInfo, len(row))
	for i, a := range row {
		addInfo[i] = dto.AddInfo{
			Key:      a.Key,
			Value:    a.Value,
			DataType: a.DataType,
			Name:     a.Name,
		}
	}
	return addInfo
}
