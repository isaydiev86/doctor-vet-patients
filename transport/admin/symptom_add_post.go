package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// SymptomAddHandler Создать новый симптом
//
//	@Summary		Создать новый симптом
//	@Description	Создать новый симптом
//	@ID				create_symptom
//	@Tags			symptoms
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.NameAdd	true	"Запрос"
//
//	@Success		200		{object}	models.Response	"Успешный ответ"
//	@Failure		400		{object}	models.Response	"Ошибка запроса"
//	@Failure		500		{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/symptoms [post]
func (s *Server) SymptomAddHandler(c *fiber.Ctx) error {
	symptom, ok := c.Locals("parsedRequest").(*models.NameAdd)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	err := s.svc.CreateSymptom(c.Context(), symptom.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:        fiber.StatusOK,
		Message:     "Success create symptom",
		Description: "Success create symptom",
	})
}
