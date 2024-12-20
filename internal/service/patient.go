package service

import (
	"context"
	"fmt"

	"github.com/isaydiev86/doctor-vet-patients/db"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/pkg/errors"
)

func (s *Service) CreatePatient(ctx context.Context, patient dto.Patient) (int64, error) {
	var patientID int64

	err := s.svc.DB.Tx(ctx, func(tx any) error {
		txDB, ok := tx.(*db.DB)
		if !ok {
			return errors.New("failed to cast transaction to *dbutil.DB")
		}

		var err error
		// Создаем пациента
		patientID, err = txDB.CreatePatient(ctx, patient)
		if err != nil {
			return fmt.Errorf("failed to create patient: %w", err)
		}

		// Создаем лечение для пациента
		_, err = txDB.CreateTreatment(ctx, patientID)
		if err != nil {
			return fmt.Errorf("failed to create treatment: %w", err)
		}

		return nil
	})

	return patientID, err
}

func (s *Service) UpdatePatient(ctx context.Context, patient dto.Patient) error {
	return s.svc.DB.UpdatePatient(ctx, patient)
}
