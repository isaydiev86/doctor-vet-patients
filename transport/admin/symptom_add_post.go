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
	var symptom models.NameAdd
	if err := c.BodyParser(&symptom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	if err := validate.Struct(symptom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Validation failed",
			Description: err.Error(),
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
		Message:     "Success",
		Description: "Success",
	})
}
