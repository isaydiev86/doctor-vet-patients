package service

import (
	"context"

	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
)

func (s *Service) GetReferences(ctx context.Context, typeQuery string) ([]*dto.Reference, error) {
	return s.svc.DB.GetReferences(ctx, typeQuery)
}

func (s *Service) GetSymptoms(ctx context.Context) ([]dto.Symptoms, error) {
	return s.svc.DB.GetSymptoms(ctx)
}

func (s *Service) GetPreparations(ctx context.Context) ([]dto.Preparations, error) {
	return s.svc.DB.GetPreparations(ctx)
}

func (s *Service) GetPreparationsToSymptoms(ctx context.Context, ids []int64) ([]dto.PreparationsWithSimilar, error) {
	// получаю список популярных препаратов по категориям
	preparations, err := s.svc.DB.GetPreparationsToSymptoms(ctx, ids)
	if err != nil {
		return nil, err
	}

	preparationsAll, err := s.svc.DB.GetPreparations(ctx)
	if err != nil {
		return nil, err
	}

	// Подсчитываем количество уникальных категорий
	uniqueCategories := make(map[string]struct{})
	for _, prep := range preparationsAll {
		uniqueCategories[prep.Category] = struct{}{}
	}

	prMap := make(map[string][]dto.Preparations, len(uniqueCategories))
	for _, prep := range preparationsAll {
		prMap[prep.Category] = append(prMap[prep.Category], prep)
	}

	preparationsWithSimilar := make([]dto.PreparationsWithSimilar, len(preparations))
	for i, pr := range preparations {
		preparationsWithSimilar[i] = dto.PreparationsWithSimilar{
			Preparations: pr,
			Similar:      mapToSimilar(prMap[pr.Category]),
		}
	}

	return preparationsWithSimilar, nil
}

func mapToSimilar(list []dto.Preparations) []dto.NameResponse {
	similar := make([]dto.NameResponse, len(list))

	for i, p := range list {
		similar[i] = dto.NameResponse{ID: p.ID, Name: p.Name}
	}

	return similar
}
