package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

type Database interface {
	GetPatients(ctx context.Context) ([]dto.Patient, error)
	CreatePatient(ctx context.Context, patient dto.Patient) error
}
