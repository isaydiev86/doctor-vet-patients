package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// TreatmentUpdateStatusHandler Обновление статуса лечения
//
//	@Summary		Обновление статуса лечения
//	@Description	Обновление статуса лечения
//	@ID				update_status_treatment
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.TreatmentUpdateStatus	true	"Запрос"
//
//	@Success		200		{object}	models.Response					"Успешный ответ"
//	@Failure		400		{object}	models.Response					"Ошибка запроса"
//	@Failure		500		{object}	models.Response					"Внутренняя ошибка сервера"
//	@Router			/treatmentUpdateStatus [put]
func (s *Server) TreatmentUpdateStatusHandler(c *fiber.Ctx) error {
	treatment, ok := c.Locals("parsedRequest").(*models.TreatmentUpdateStatus)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	err := s.svc.UpdateStatusTreatment(c.Context(), mapDtoUpdateStatusOfApi(treatment))
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

func mapDtoUpdateStatusOfApi(api *models.TreatmentUpdateStatus) dto.TreatmentUpdateStatus {
	return dto.TreatmentUpdateStatus{
		ID:     api.ID,
		Status: api.Status,
	}
}
