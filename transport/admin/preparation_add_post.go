package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// PreparationAddHandler Создать новый препарат
//
//	@Summary		Создать новый препарат
//	@Description	Создать новый препарат
//	@ID				create_preparation
//	@Tags			preparations
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.PreparationsAdd	true	"Запрос"
//
//	@Success		200		{object}	models.Response			"Успешный ответ"
//	@Failure		400		{object}	models.Response			"Ошибка запроса"
//	@Failure		500		{object}	models.Response			"Внутренняя ошибка сервера"
//	@Router			/preparations [post]
func (s *Server) PreparationAddHandler(c *fiber.Ctx) error {
	pr, ok := c.Locals("parsedRequest").(*models.PreparationsAdd)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	err := s.svc.CreatePreparations(c.Context(), mapPreparationDtoOfApi(pr))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:        fiber.StatusOK,
		Message:     "Success create preparations",
		Description: "Success create preparations",
	})
}

func mapPreparationDtoOfApi(add *models.PreparationsAdd) dto.PreparationsAdd {
	return dto.PreparationsAdd{
		Name:     add.Name,
		Dose:     add.Dose,
		Course:   add.Course,
		Category: add.Category,
		Option:   add.Option,
	}
}
