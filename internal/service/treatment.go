package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/db"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error) {
	return s.svc.DB.GetTreatments(ctx, filter)
}

func (s *Service) GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error) {
	return s.svc.DB.GetTreatment(ctx, id)
}

func (s *Service) GetTreatmentForUser(ctx context.Context, userId string) (*dto.TreatmentDetail, error) {
	return s.svc.DB.GetTreatmentForUser(ctx, userId)
}

func (s *Service) UpdateTreatmentForUser(ctx context.Context, treatment dto.TreatmentSendForUser) error {
	return s.svc.DB.UpdateTreatmentForUser(ctx, treatment)
}

func (s *Service) UpdateTreatment(ctx context.Context, treatment dto.TreatmentUpdateToUser) error {
	return s.svc.DB.Tx(ctx, func(tx any) error {
		txDB, ok := tx.(*db.DB)
		if !ok {
			return errors.New("failed to cast transaction to *dbutil.DB")
		}

		err := txDB.UpdateTreatment(ctx, treatment)
		if err != nil {
			return fmt.Errorf("failed db UpdateTreatment: %w", err)
		}

		// Add prescriptions for treatment
		err = txDB.AddPrescriptionsToTreatment(ctx, treatment.ID, treatment.Prescriptions)
		if err != nil {
			return fmt.Errorf("failed db AddPrescriptionsToTreatment: %w", err)
		}

		///TODO update popularity for preparation with PreparationID

		return nil
	})
}
