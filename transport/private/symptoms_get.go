package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// SymptomsHandler Симптомы
//
//	@Summary		Получить список симптомов
//	@Description	Получить список симптомов
//	@ID				get_symptoms
//	@Tags			symptoms
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Symptoms	"Список симптомов"
//	@Failure		400	{object}	models.Response	"Ошибка запроса"
//	@Failure		500	{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/symptoms [get]
func (s *Server) SymptomsHandler(c *fiber.Ctx) error {
	symptomsData, err := s.svc.GetSymptoms(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	symptoms := make([]models.Symptoms, len(symptomsData))
	for i, p := range symptomsData {
		symptoms[i] = models.Symptoms{
			ID:   p.ID,
			Name: p.Name,
		}
	}

	return c.JSON(symptoms)
}
