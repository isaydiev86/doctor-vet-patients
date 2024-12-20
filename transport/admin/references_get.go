package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// ReferenceHandler Справочник
//
//	@Summary		Получить список справочника
//	@Description	Получить список справочника
//	@ID				get_reference
//	@Tags			reference
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string				true	"тип справочника"
//	@Success		200		{array}		models.RefResponse	"Список справочника"
//	@Failure		400		{object}	models.Response		"Ошибка запроса"
//	@Failure		500		{object}	models.Response		"Внутренняя ошибка сервера"
//	@Router			/reference [get]
func (s *Server) ReferenceHandler(c *fiber.Ctx) error {
	typeQuery := c.Query("type", "symptoms")

	referenceData, err := s.svc.GetReferences(c.Context(), typeQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	references := make([]models.RefResponse, len(referenceData))
	for i, p := range referenceData {
		references[i] = models.RefResponse{
			ID:   p.ID,
			Name: p.Name,
			Type: p.Type,
		}
	}

	return c.JSON(references)
}
