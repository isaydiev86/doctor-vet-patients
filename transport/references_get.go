package transport

import (
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport/models"
	"github.com/gofiber/fiber/v2"
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
func ReferenceHandler(c *fiber.Ctx, svc service.Service) error {
	typeQuery := c.Query("type", "symptoms")

	referenceData, err := svc.GetReferences(c.Context(), typeQuery)
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
