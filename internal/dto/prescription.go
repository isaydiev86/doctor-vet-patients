package dto

import "time"

type Prescription struct {
	ID          int64
	TreatmentID int64
	Name        string
	Dose        float64
	Course      string
	Category    string
	Option      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
