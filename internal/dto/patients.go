package dto

type Patient struct {
	ID         int64
	DoctorID   string
	Fio        string
	Phone      string
	Address    string
	Animal     string
	Name       string
	Breed      string
	Age        float64
	Gender     string
	Status     string
	IsNeutered bool
}
