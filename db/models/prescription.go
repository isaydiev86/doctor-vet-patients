package models

import (
	"database/sql"
	"time"
)

// Prescription лечение, которое оказали пациенту
type Prescription struct {
	ID          int64          `db:"id" json:"id"`
	TreatmentID int64          `db:"treatment_id" json:"treatment_id"`
	Preparation sql.NullString `db:"preparation" json:"preparation"`
	Dose        sql.NullString `db:"dose" json:"dose"`
	Course      sql.NullString `db:"course" json:"course"`
	Amount      sql.NullString `db:"amount" json:"amount"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
