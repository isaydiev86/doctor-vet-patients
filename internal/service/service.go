package service

type Relation struct {
	DB Database
}

func New(svc Relation) *Service {
	return &Service{svc: svc}
}

type Service struct {
	svc Relation
}
