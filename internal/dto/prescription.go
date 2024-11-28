package dto

import "time"

type Prescription struct {
	ID          int64
	TreatmentID int64
	Preparation string
	Dose        string
	Course      string
	Amount      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
