package dto

type Service struct {
	ID          int64
	TreatmentID int64
	DoctorID    string
	Name        string
	Price       float64
}

type ServiceRef struct {
	ID    int64
	Name  string
	Price float64
}

// Other расходники
type Other struct {
	ID     int64
	Amount int64
	Name   string
}
