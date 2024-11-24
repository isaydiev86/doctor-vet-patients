package db

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

type IStorage interface {
	GetPatients(ctx context.Context) ([]dto.Patient, error)
	CreatePatient(ctx context.Context, patient dto.Patient) error
}
