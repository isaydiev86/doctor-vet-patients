package service

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (s *Service) GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error) {
	list, err := s.svc.DB.GetTreatments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error) {
	return s.svc.DB.GetTreatment(ctx, id)
}

//func (s *Service) Tx(ctx context.Context, f func(any) error) error {
//	return s.svc.DB.Tx(ctx, f)
//}
