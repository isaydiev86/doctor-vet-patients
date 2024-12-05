package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	return s.svc.DB.GetReferences(ctx, typeQuery)
}
