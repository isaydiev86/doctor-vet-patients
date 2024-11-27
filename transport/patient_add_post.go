package transport

import (
	"context"

	"doctor-vet-patients/mapper"
	"doctor-vet-patients/transport/models"
	"github.com/gofiber/fiber/v2"
)

// PatientAddHandler Получить список пациентов
//
//	@Summary		Создать нового пациента
//	@Description	Создать нового пациента
//	@ID				create_patient
//	@Tags			patients
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.Patient	true	"Запрос"
//
//	@Success		200		{object}	models.Response	"Успешный ответ"
//	@Failure		400		{object}	models.Response	"Ошибка запроса"
//	@Failure		500		{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/create_patient [post]
func PatientAddHandler(c *fiber.Ctx, svc service.IService) error {
	ctx := context.Background()

	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	// Получаем пациентов через сервис
	err := svc.CreatePatient(ctx, mapper.GetDtoPatientOfApi(patient))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Error",
			Description: "Error for backend",
		})
	}

	return c.JSON(models.Response{
		Code:        200,
		Message:     "Success",
		Description: "Success",
	})
}
