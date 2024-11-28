package transport

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

type Services interface {
	GetTreatments(ctx context.Context) ([]dto.Patient, error)
	CreatePatient(ctx context.Context, patient dto.Patient) error
}
