package models

import "time"

// Treatment лечение пациента (заявки)
type Treatment struct {
	ID          int64      `json:"id" validate:"required"`
	PatientID   int64      `json:"patientId" validate:"required"`
	DoctorID    string     `json:"doctorId" validate:"required"`
	Status      string     `json:"status" validate:"required"`
	CreatedAt   time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt" validate:"required"`
	BeginAt     *time.Time `json:"beginAt" validate:"required"`
	EndAt       *time.Time `json:"endAt" validate:"required"`
	Comment     string     `json:"comment" validate:"required"`
	IsActive    int64      `json:"isActive" validate:"required"`
	Weight      float64    `json:"weight" validate:"gte=0"`      // Вес животного в кг, >= 0
	Temperature float64    `json:"temperature" validate:"gte=0"` // Температура тела животного, >= 0

	Patient Patient `json:"patient"` // инфа пациента
}

// TreatmentDetail детали лечение
type TreatmentDetail struct {
	ID          int64      `json:"id" validate:"required"`
	PatientID   int64      `json:"patientId" validate:"required"`
	DoctorID    string     `json:"doctorId" validate:"required"`
	Status      string     `json:"status" validate:"required"`
	CreatedAt   time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time  `json:"updatedAt" validate:"required"`
	BeginAt     *time.Time `json:"beginAt" validate:"required"`
	EndAt       *time.Time `json:"endAt" validate:"required"`
	Comment     string     `json:"comment" validate:"required"`
	IsActive    int64      `json:"isActive" validate:"required"`
	Weight      float64    `json:"weight" validate:"gte=0"`      // Вес животного в кг, >= 0
	Temperature float64    `json:"temperature" validate:"gte=0"` // Температура тела животного, >= 0

	Patient       Patient        `json:"patient"`       // инфа пациента
	Prescriptions []Prescription `json:"prescriptions"` // список лечения

	AddInfo []AddInfo `json:"addInfo"` // список доп полей
}

// TreatmentSendForUser реквест для передачи лечение на доктора
type TreatmentSendForUser struct {
	ID       int64  `json:"id" validate:"required"`
	DoctorID string `json:"doctorId" validate:"required"`
}

// TreatmentUpdateStatus Обновление статуса лечения
type TreatmentUpdateStatus struct {
	ID     int64  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required,oneof=процесс завершен отклонен ожидает закрыта"`
}

// TreatmentUpdateToUser обновление лечения доктором
type TreatmentUpdateToUser struct {
	ID          int64   `json:"id" validate:"required"`
	DoctorID    string  `json:"doctorId" validate:"required"`
	Weight      float64 `json:"weight" validate:"gte=0"`
	Temperature float64 `json:"temperature" validate:"gte=0"`
	Comment     string  `json:"comment"`

	Prescriptions []PrescriptionForUpdate `json:"prescriptions" validate:"min=1"`

	AddInfo []AddInfo `json:"addInfo"`
}

type AddInfo struct {
	Key      string      `json:"key" validate:"required"`
	Value    interface{} `json:"value"`
	DataType string      `json:"dataType"`
	Name     string      `json:"name"`
}
