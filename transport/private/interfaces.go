package private

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
	GetTreatmentForUser(ctx context.Context, userId string) (*dto.TreatmentDetail, error)
	UpdateTreatment(ctx context.Context, treatment dto.TreatmentUpdateToUser) error

	GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error)
	GetSymptoms(ctx context.Context) ([]*dto.Symptoms, error)
	GetPreparations(ctx context.Context) ([]*dto.Preparations, error)
	GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]*dto.Preparations, error)
}
