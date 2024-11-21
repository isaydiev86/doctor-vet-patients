package transport

import (
	"context"

	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport/models"
	"github.com/gofiber/fiber/v2"
)

// PatientHandler Получить список пациентов
//
//	@Summary		Получить список пациентов
//	@Description	Возвращает список всех пациентов
//	@ID				get_patients
//	@Tags			patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Patient			"Список пациентов"
//	@Failure		400	{object}	models.ErrorResponse	"Ошибка запроса"
//	@Failure		500	{object}	models.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/patients [get]
func PatientHandler(c *fiber.Ctx, svc service.IService) error {
	ctx := context.Background()

	// Получаем пациентов через сервис
	patients, err := svc.GetPatients(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Code:        fiber.StatusInternalServerError,
			Message:     "Error",
			Description: "Error for backend",
		})
	}

	return c.JSON(patients)
}
