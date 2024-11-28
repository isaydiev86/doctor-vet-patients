package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (s *Service) CreatePatient(ctx context.Context, patient dto.Patient) error {
	return s.svc.DB.CreatePatient(ctx, patient)
}
