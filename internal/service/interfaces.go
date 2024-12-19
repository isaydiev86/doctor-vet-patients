package service

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
}

type Database interface {
	Tx(ctx context.Context, f func(any) error) error

	GetTreatments(ctx context.Context, filter dto.TreatmentFilters) ([]*dto.Treatment, error)
	GetTreatment(ctx context.Context, id int64) (*dto.TreatmentDetail, error)
	CreateTreatment(ctx context.Context, patientID int64) (int64, error)

	CreatePatient(ctx context.Context, patient dto.Patient) (int64, error)
	UpdatePatient(ctx context.Context, patient dto.Patient) error

	GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error)

	UserExists(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, userID, name, role string) error
	GetUsers(ctx context.Context, filter dto.UserFilters) ([]*dto.User, error)
}
