package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
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
//	@Param			Form	body		models.PatientUpdate	true	"Запрос"
//
//	@Success		200		{object}	models.Response	"Успешный ответ"
//	@Failure		400		{object}	models.Response	"Ошибка запроса"
//	@Failure		500		{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/patient [put]
func (s *Server) PatientUpdateHandler(c *fiber.Ctx) error {
	patient, ok := c.Locals("parsedRequest").(*models.PatientUpdate)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	// Получаем пациентов через сервис
	err := s.svc.UpdatePatient(c.Context(), getDtoUpdatePatientOfApi(patient))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	return c.JSON(models.Response{
		Code:        fiber.StatusOK,
		Message:     "Success update patient",
		Description: "Success update patient",
	})
}

func getDtoUpdatePatientOfApi(patient *models.PatientUpdate) dto.Patient {
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
