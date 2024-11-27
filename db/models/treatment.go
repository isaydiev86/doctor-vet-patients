package models

import "time"

type StatusPrescription string

const (
	InProcess StatusPrescription = "в процессе" // когда отдали конкретному врачу
	Done      StatusPrescription = "завершен"
	Decline   StatusPrescription = "отклонен"
	InLine    StatusPrescription = "в очереди" // когда новая заявка
)

func (s StatusPrescription) String() string { return string(s) }

// Treatment лечение пациента
type Treatment struct {
	ID          int64   `db:"id"`
	PatientID   int64   `db:"patient_id"`
	DoctorID    string  `db:"doctor_id"`
	Temperature float64 `db:"temperature"`
	Status      StatusPrescription
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	BeginAt     time.Time `db:"begin_at"`  // начало лечения
	EndAt       time.Time `db:"end_at"`    //конец лечения
	comment     string    `db:"comment"`   // информация в случаи закрытия лечения
	isActive    int64     `db:"is_active"` // завершено лечение или нет 1 - нет 0 - да
	Age         float64   `db:"age"`
	Weight      float64   `db:"weight"`
}
