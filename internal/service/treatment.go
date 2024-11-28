package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (s *Service) GetTreatments(ctx context.Context) ([]dto.Treatment, error) {
	return s.svc.DB.GetTreatments(ctx)
}

func (s *Service) GetTreatment(ctx context.Context, id int64) (dto.TreatmentDetail, error) {
	return s.svc.DB.GetTreatment(ctx, id)
}
