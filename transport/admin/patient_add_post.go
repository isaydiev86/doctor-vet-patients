package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// PatientAddHandler Создать нового пациента
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
//	@Router			/patient [post]
func (s *Server) PatientAddHandler(c *fiber.Ctx) error {
	patient, ok := c.Locals("parsedRequest").(*models.Patient)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
		})
	}

	_, err := s.svc.CreatePatient(c.Context(), getDtoPatientOfApi(patient))
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

func getDtoPatientOfApi(patient *models.Patient) dto.Patient {
	return dto.Patient{
		Fio:        patient.Fio,
		Phone:      patient.Phone,
		Animal:     utils.FromPtr(patient.Animal),
		Name:       patient.Name,
		Breed:      utils.FromPtr(patient.Breed),
		Age:        utils.FromPtr(patient.Age),
		Gender:     patient.Gender,
		IsNeutered: patient.IsNeutered,
	}
}
