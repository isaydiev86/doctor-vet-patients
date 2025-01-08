package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// RelationSymptomWithPreparationHandler Создать связь симптома с препаратом
//
//	@Summary		Создать связь симптома с препаратом
//	@Description	Создать связь симптома с препаратом
//	@ID				relation_symptom_with_preparation
//	@Tags			relations
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.RelationSymptomWithPreparation	true	"Запрос"
//
//	@Success		200		{object}	models.Response							"Успешный ответ"
//	@Failure		400		{object}	models.Response							"Ошибка запроса"
//	@Failure		500		{object}	models.Response							"Внутренняя ошибка сервера"
//	@Router			/relationSymptomWithPreparation [post]
func (s *Server) RelationSymptomWithPreparationHandler(c *fiber.Ctx) error {
	relation, ok := c.Locals("parsedRequest").(*models.RelationSymptomWithPreparation)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}
	err := s.svc.AddRelationSymptomWithPreparation(c.Context(), relation.SymptomID, relation.PreparationID)
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
