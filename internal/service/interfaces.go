package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

type Database interface {
	Tx(ctx context.Context, f func(any) error) error

	GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error)
	GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error)
	CreateTreatment(ctx context.Context, patientID int64) (int64, error)

	CreatePatient(ctx context.Context, patient dto.Patient) (int64, error)
}
