package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type StatusPrescription string

const (
	InProcess StatusPrescription = "процесс" // когда отдали конкретному врачу
	Done      StatusPrescription = "завершен"
	Decline   StatusPrescription = "отклонен"
	Wait      StatusPrescription = "ожидает" // когда новая заявка
	End       StatusPrescription = "закрыта" // закрытие лечения
)

func (s StatusPrescription) String() string { return string(s) }

// Treatment лечение пациента
type Treatment struct {
	ID          int64           `db:"id"`
	PatientID   int64           `db:"patient_id"`
	IsActive    int64           `db:"is_active"` // завершено лечение или нет 1 - нет 0 - да
	DoctorID    string          `db:"doctor_id"`
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at"`
	Comment     sql.NullString  `db:"comment"` // информация в случаи закрытия лечения
	Status      sql.NullString  `db:"status"`
	Weight      sql.NullFloat64 `db:"weight"`
	Temperature sql.NullFloat64 `db:"temperature"`
	BeginAt     sql.NullTime    `db:"begin_at"` // начало лечения
	EndAt       sql.NullTime    `db:"end_at"`   // конец лечения
}

type TreatmentRow struct {
	ID          int64           `db:"id"`
	PatientID   int64           `db:"patient_id"`
	DoctorID    sql.NullString  `db:"doctor_id"`
	Status      sql.NullString  `db:"status"`
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at"`
	BeginAt     sql.NullTime    `db:"begin_at"`  // начало лечения
	EndAt       sql.NullTime    `db:"end_at"`    // конец лечения
	Comment     sql.NullString  `db:"comment"`   // информация в случаи закрытия лечения
	IsActive    int64           `db:"is_active"` // завершено лечение или нет 1 - нет 0 - да
	Weight      sql.NullFloat64 `db:"weight"`
	Temperature sql.NullFloat64 `db:"temperature"`
	Patient     `db:"patient"`
}

type TreatmentDetailRow struct {
	ID                int64           `db:"id"`
	PatientID         int64           `db:"patient_id"`
	DoctorID          sql.NullString  `db:"doctor_id"`
	Status            sql.NullString  `db:"status"`
	CreatedAt         time.Time       `db:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at"`
	BeginAt           sql.NullTime    `db:"begin_at"`  // начало лечения
	EndAt             sql.NullTime    `db:"end_at"`    // конец лечения
	Comment           sql.NullString  `db:"comment"`   // информация в случаи закрытия лечения
	IsActive          int64           `db:"is_active"` // завершено лечение или нет 1 - нет 0 - да
	Weight            sql.NullFloat64 `db:"weight"`
	Temperature       sql.NullFloat64 `db:"temperature"`
	Patient           `db:"patient"`
	PrescriptionsJSON json.RawMessage `db:"prescriptions"`
	Prescription      []Prescription  `db:"prescription" json:"prescription"`
}
