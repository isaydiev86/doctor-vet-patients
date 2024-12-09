package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetUsers(ctx context.Context, filter dto.UserFilters) ([]*dto.User, error) {
	return s.svc.DB.GetUsers(ctx, filter)
}
