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

//type IService interface {
//	GetPatients(ctx context.Context) ([]dto.Patient, error)
//	CreatePatient(ctx context.Context, patient dto.Patient) error
//}
//
//type service struct {
//	storage db.IStorage
//}
//
//func NewService(storage db.IStorage) IService {
//	return &service{storage: storage}
//}
//
//func (s *service) CreatePatient(ctx context.Context, patient dto.Patient) error {
//
//	/// TODO create Prescription(первоначальное назначение - без лечения)
//
//	return s.storage.CreatePatient(ctx, patient)
//}
//
//func (s *service) GetPatients(ctx context.Context) ([]dto.Patient, error) {
//	return s.storage.GetPatients(ctx)
//}
