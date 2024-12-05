package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// PatientUpdateHandler Редактирование  пациента
//
//	@Summary		Редактирование  пациента
//	@Description	Редактирование  пациента
//	@ID				update_patient
//	@Tags			patients
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.Patient	true	"Запрос"
//
//	@Success		200		{object}	models.Response	"Успешный ответ"
//	@Failure		400		{object}	models.Response	"Ошибка запроса"
//	@Failure		500		{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/patient [put]
func PatientUpdateHandler(c *fiber.Ctx, svc service.Service) error {
	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	if err := validate.Struct(patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Validation failed",
			Description: err.Error(),
		})
	}

	// Получаем пациентов через сервис
	err := svc.UpdatePatient(c.Context(), getDtoUpdatePatientOfApi(patient))
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

func getDtoUpdatePatientOfApi(patient models.Patient) dto.Patient {
	return dto.Patient{
		ID:         patient.ID,
		Fio:        patient.Fio,
		Phone:      patient.Phone,
		Address:    patient.Address,
		Animal:     patient.Animal,
		Name:       patient.Name,
		Breed:      patient.Breed,
		Age:        patient.Age,
		Gender:     patient.Gender,
		IsNeutered: patient.IsNeutered,
	}
}
