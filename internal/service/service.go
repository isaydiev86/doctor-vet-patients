package service

import (
	"go.uber.org/zap"
)

type Relation struct {
	DB Database
}

func New(svc Relation, logger *zap.Logger, keycloak *KeycloakService) *Service {
	return &Service{svc: svc, Logger: logger, Keycloak: keycloak}
}

type Service struct {
	svc      Relation
	Logger   *zap.Logger
	Keycloak *KeycloakService
}
