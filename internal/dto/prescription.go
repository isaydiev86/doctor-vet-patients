package dto

type Prescription struct {
	ID          int64
	TreatmentID int64
	Preparation string
	Dose        string
	Course      string
	Amount      string
}
