package transport

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

type Services interface {
	Login(ctx context.Context, login dto.LoginRequest) (*dto.LoginResponse, error)

	GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error)
	GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error)

	CreatePatient(ctx context.Context, patient dto.Patient) (int64, error)
	UpdatePatient(ctx context.Context, patient dto.Patient) error

	GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error)
}
