package entity

import "time"

type StatusPrescription string

const (
	InProcess StatusPrescription = "в процессе" // когда отдали конкретному врачу
	Done      StatusPrescription = "завершен"
	Decline   StatusPrescription = "отклонен"
	InLine    StatusPrescription = "в очереди" // когда новая заявка
)

func (s StatusPrescription) String() string { return string(s) }

// Patient Основная инфа пациента
type Patient struct {
	ID         int64
	Fio        string
	Phone      string
	Address    string
	Animal     string
	Name       string
	Breed      string
	Gender     string
	IsNeutered bool

	Prescriptions []InitialPrescription
}

// InitialPrescription первоначальные данные пациента
type InitialPrescription struct {
	ID          int64
	PatientID   int64
	DoctorID    string
	Temperature float64
	Status      StatusPrescription
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`

	OpenAt   time.Time // начало лечения
	ClosedAt time.Time //конец лечения
	comment  string    // информация в случаи закрытия лечения
	isActive int64     // завершено лечение или нет 1 - нет 0 - да

	Prescriptions []Prescription

	Age    float64
	Weight float64
}

// Prescription лечение, которое оказали пациенту
type Prescription struct {
	ID                    int64
	InitialPrescriptionID int64
	Preparation           string
	Dose                  string
	Course                string
	Amount                string
}

// Doctor Основная инфа врача
type Doctor struct {
	ID    string // айди из кейклоака
	Fio   string
	Phone string
	Role  string // роль задачется в кейклоаке
}

// History История лечения конкретного пациента - одна запись за день
type History struct {
	ID                    int64
	InitialPrescriptionID int64
	comment               string    // комент от врача
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}
