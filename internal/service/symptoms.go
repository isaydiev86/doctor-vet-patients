package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetSymptoms(ctx context.Context) ([]dto.Symptoms, error) {
	return s.svc.DB.GetSymptoms(ctx)
}

func (s *Service) CreateSymptom(ctx context.Context, name string) error {
	return s.svc.DB.CreateSymptom(ctx, name)
}
