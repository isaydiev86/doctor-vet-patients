package models

import "time"

// History История лечения конкретного пациента - одна запись за день
type History struct {
	ID          int64     `db:"id"`
	TreatmentID int64     `db:"treatment_iD"`
	Comment     string    `db:"comment"` // комент от врача
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
