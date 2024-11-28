package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (s *Service) GetTreatments(ctx context.Context) ([]dto.Treatment, error) {
	return s.svc.DB.GetTreatments(ctx)
}
