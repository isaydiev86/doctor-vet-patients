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
	BeginAt     time.Time
	EndAt       time.Time
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
	BeginAt     time.Time
	EndAt       time.Time
	Comment     string
	IsActive    int64
	Temperature float64
	Weight      float64

	Patient      Patient // поля самого пациента потом мб сделаем отдельную dto, где меньше полей
	Prescription []Prescription
}

type TreatmentFilters struct {
	Fio    string
	Name   string
	Status string
	Limit  int
	Offset int
}
