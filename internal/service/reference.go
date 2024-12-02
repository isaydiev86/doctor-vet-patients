package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (s *Service) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	return s.svc.DB.GetReferences(ctx, typeQuery)
}
