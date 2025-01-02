package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Prescription лечение, которое оказали пациенту
type Prescription struct {
	ID          int64           `db:"id" json:"id"`
	TreatmentID int64           `db:"treatment_id" json:"treatment_id"`
	Name        sql.NullString  `db:"name" json:"name"`
	Dose        sql.NullFloat64 `db:"dose" json:"dose"`
	Course      sql.NullString  `db:"course" json:"course"`
	Category    sql.NullString  `db:"category" json:"category"`
	Option      sql.NullString  `db:"option" json:"option"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

// PrescriptionRow лечение, которое оказали пациенту для update
type PrescriptionRow struct {
	TreatmentID int64           `db:"treatment_id" json:"treatment_id"`
	Name        sql.NullString  `db:"name" json:"name"`
	Dose        sql.NullFloat64 `db:"dose" json:"dose"`
	Course      sql.NullString  `db:"course" json:"course"`
	Category    sql.NullString  `db:"category" json:"category"`
	Option      sql.NullString  `db:"option" json:"option"`
}

func (p *Prescription) UnmarshalJSON(data []byte) error {
	type Alias Prescription
	aux := &struct {
		Name     string  `json:"name"`
		Dose     float64 `json:"dose"`
		Course   string  `json:"course"`
		Category string  `json:"category"`
		Option   string  `json:"option"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Преобразуем обычные значения в sql.NullString и sql.NullFloat64
	p.Name = sql.NullString{String: aux.Name, Valid: aux.Name != ""}
	p.Dose = sql.NullFloat64{Float64: aux.Dose, Valid: aux.Dose != 0}
	p.Course = sql.NullString{String: aux.Course, Valid: aux.Course != ""}
	p.Category = sql.NullString{String: aux.Category, Valid: aux.Category != ""}
	p.Option = sql.NullString{String: aux.Option, Valid: aux.Option != ""}

	return nil
}
