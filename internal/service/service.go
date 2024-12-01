package service

import "go.uber.org/zap"

type Relation struct {
	DB Database
}

func New(svc Relation, logger *zap.Logger) *Service {
	return &Service{svc: svc, Logger: logger}
}

type Service struct {
	svc    Relation
	Logger *zap.Logger
}
