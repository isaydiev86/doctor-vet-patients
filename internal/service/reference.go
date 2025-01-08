package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	return s.svc.DB.GetReferences(ctx, typeQuery)
}

func (s *Service) AddRelationSymptomWithPreparation(ctx context.Context, symptomID, preparationID int64) error {
	return s.svc.DB.AddRelationSymptomWithPreparation(ctx, symptomID, preparationID)
}
