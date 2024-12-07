package service

import (
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"go.uber.org/zap"
)

type Relation struct {
	DB Database
}

func New(svc Relation, logger *zap.Logger, keycloak *keycloak.Service) *Service {
	return &Service{svc: svc, Logger: logger, Keycloak: keycloak}
}

type Service struct {
	svc      Relation
	Logger   *zap.Logger
	Keycloak *keycloak.Service
}
