package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	return s.svc.DB.GetReferences(ctx, typeQuery)
}

func (s *Service) GetSymptoms(ctx context.Context) ([]*dto.Symptoms, error) {
	return s.svc.DB.GetSymptoms(ctx)
}

func (s *Service) GetPreparations(ctx context.Context) ([]*dto.Preparations, error) {
	return s.svc.DB.GetPreparations(ctx)
}

func (s *Service) GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]*dto.Preparations, error) {
	return s.svc.DB.GetPreparationsToSymptoms(ctx, ids)
}
