package dto

import "time"

// Treatment лечение пациента (заявки)
type Treatment struct {
	ID          int64
	PatientID   int64
	DoctorID    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BeginAt     *time.Time
	EndAt       *time.Time
	Comment     string
	IsActive    int64
	Weight      float64
	Temperature float64

	Patient Patient // поля самого пациента потом мб сделаем отдельную dto, где меньше полей
}

// TreatmentDetail детали лечения
type TreatmentDetail struct {
	ID          int64
	PatientID   int64
	DoctorID    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BeginAt     *time.Time
	EndAt       *time.Time
	Comment     string
	IsActive    int64
	Temperature float64
	Weight      float64

	Patient      Patient // поля самого пациента потом мб сделаем отдельную dto, где меньше полей
	Prescription []Prescription
	AddInfo      []AddInfo

	//	добавить список оказанных услуг services []Service
	//	добавить расходники services []Others
	// добавить поле isSelf - если пациент со своими препаратами(чтобы не вычитать из склада)

}

type TreatmentFilters struct {
	Fio    string
	Name   string
	Status string
	Date   string
	Limit  int
	Offset int
}

type TreatmentSendForUser struct {
	ID       int64
	DoctorID string
}

type TreatmentUpdateStatus struct {
	ID     int64
	Status string
}

type TreatmentUpdateToUser struct {
	ID          int64
	DoctorID    string
	Weight      float64
	Temperature float64
	Comment     string

	Prescriptions []PrescriptionForUpdate
	AddInfo       []AddInfo
}

type AddInfo struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	DataType string      `json:"dataType"`
	Name     string      `json:"name"`
}
