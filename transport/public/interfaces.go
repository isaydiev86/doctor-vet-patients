package public

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
	Login(ctx context.Context, login dto.LoginRequest) (*dto.LoginResponse, error)
}
