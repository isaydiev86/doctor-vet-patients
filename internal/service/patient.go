package service

import (
	"context"
	"fmt"

	"doctor-vet-patients/db"
	"doctor-vet-patients/internal/dto"
	"github.com/pkg/errors"
)

func (s *Service) CreatePatient(ctx context.Context, patient dto.Patient) error {
	return s.svc.DB.Tx(ctx, func(tx any) error {
		txDB, ok := tx.(*db.DB)
		if !ok {
			return errors.New("failed to cast transaction to *dbutil.DB")
		}

		// Создаем пациента
		patientID, err := txDB.CreatePatient(ctx, patient)
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
}

func (s *Service) UpdatePatient(ctx context.Context, patient dto.Patient) error {
	return s.svc.DB.UpdatePatient(ctx, patient)
}
