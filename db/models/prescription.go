package models

// Prescription лечение, которое оказали пациенту
type Prescription struct {
	ID          int64  `db:"id"`
	TreatmentID int64  `db:"treatment_id"`
	Preparation string `db:"preparation"`
	Dose        string `db:"dose"`
	Course      string `db:"course"`
	Amount      string `db:"amount"`
}
