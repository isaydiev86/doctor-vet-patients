package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// TreatmentSendOnUserHandler Назначить лечения на доктора
//
//	@Summary		Назначить лечения на доктора
//	@Description	Назначить лечения на доктора
//	@ID				treatment_send_on_user
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.TreatmentSendForUser	true	"Запрос"
//
//	@Success		200		{object}	models.Response				"Успешный ответ"
//	@Failure		400		{object}	models.Response				"Ошибка запроса"
//	@Failure		500		{object}	models.Response				"Внутренняя ошибка сервера"
//	@Router			/send_treatment [put]
func (s *Server) TreatmentSendOnUserHandler(c *fiber.Ctx) error {
	treatment, ok := c.Locals("parsedRequest").(*models.TreatmentSendForUser)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	err := s.svc.UpdateTreatmentForUser(c.Context(), dto.TreatmentSendForUser{
		ID:       treatment.ID,
		DoctorID: treatment.DoctorID,
	})
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
