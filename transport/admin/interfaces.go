package admin

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

type Logger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Fatal(msg string, args ...any)
	Sync() error
}

type Services interface {
	GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error)
	GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error)
	UpdateTreatmentForUser(ctx context.Context, treatment dto.TreatmentSendForUser) error

	CreatePatient(ctx context.Context, patient dto.Patient) (int64, error)
	UpdatePatient(ctx context.Context, patient dto.Patient) error

	GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error)

	GetUsers(ctx context.Context, role string) ([]*dto.User, error)
}
