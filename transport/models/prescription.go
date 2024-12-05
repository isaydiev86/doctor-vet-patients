package models

import "time"

type Prescription struct {
	ID          int64     `json:"id" validate:"required"`
	TreatmentID int64     `json:"treatmentId" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Dose        float64   `json:"dose" validate:"required"`
	Course      string    `json:"course" validate:"required"`
	Category    string    `json:"category"`
	Option      string    `json:"option"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}
