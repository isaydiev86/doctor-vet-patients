package models

import (
	"database/sql"
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
