package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// PreparationsHandler Препараты
//
//	@Summary		Получить список препаратов
//	@Description	Получить список препаратов
//	@ID				get_preparations
//	@Tags			preparations
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Preparations	"Список препаратов"
//	@Failure		400	{object}	models.Response		"Ошибка запроса"
//	@Failure		500	{object}	models.Response		"Внутренняя ошибка сервера"
//	@Router			/preparations [get]
func (s *Server) PreparationsHandler(c *fiber.Ctx) error {
	preparationsData, err := s.svc.GetPreparations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	preparations := make([]models.Preparations, len(preparationsData))
	for i, p := range preparationsData {
		preparations[i] = models.Preparations{
			ID:       p.ID,
			Name:     p.Name,
			Dose:     p.Dose,
			Course:   p.Course,
			Category: p.Category,
			Option:   p.Option,
		}
	}

	return c.JSON(preparations)
}
