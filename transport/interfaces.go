package transport

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

type Services interface {
	GetTreatments(ctx context.Context) ([]dto.Patient, error)
	GetTreatment(ctx context.Context, id int64) (dto.TreatmentDetail, error)
	CreatePatient(ctx context.Context, patient dto.Patient) error
}
