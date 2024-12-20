package service

import "github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"

type Relation struct {
	DB Database
}

func New(svc Relation, logger Logger, keycloak *keycloak.Service) *Service {
	return &Service{svc: svc, Logger: logger, Keycloak: keycloak}
}

type Service struct {
	svc      Relation
	Logger   Logger
	Keycloak *keycloak.Service
}
