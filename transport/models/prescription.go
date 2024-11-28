package models

import "time"

type Prescription struct {
	ID          int64     `json:"id" validate:"required"`
	TreatmentID int64     `json:"treatmentId" validate:"required"`
	Preparation string    `json:"preparation" validate:"required"`
	Dose        string    `json:"dose" validate:"required"`
	Course      string    `json:"course" validate:"required"`
	Amount      string    `json:"amount" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}
