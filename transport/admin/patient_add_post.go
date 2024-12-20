package admin

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

var validate = validator.New()

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

func getDtoPatientOfApi(patient models.Patient) dto.Patient {
	return dto.Patient{
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
