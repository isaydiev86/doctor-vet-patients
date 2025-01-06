package private

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

var validate = validator.New()

// TreatmentUpdateHandler Обновление лечения
//
//	@Summary		Обновление лечения
//	@Description	Обновление лечения
//	@ID				update_treatment
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.TreatmentUpdateToUser	true	"Запрос"
//
//	@Success		200		{object}	models.Response					"Успешный ответ"
//	@Failure		400		{object}	models.Response					"Ошибка запроса"
//	@Failure		500		{object}	models.Response					"Внутренняя ошибка сервера"
//	@Router			/treatment [put]
func (s *Server) TreatmentUpdateHandler(c *fiber.Ctx) error {
	var treatment models.TreatmentUpdateToUser
	if err := c.BodyParser(&treatment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	if err := validate.Struct(treatment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Validation failed",
			Description: err.Error(),
		})
	}

	err := s.svc.UpdateTreatment(c.Context(), mapDtoUpdateTreatmentOfApi(treatment))
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

func mapDtoUpdateTreatmentOfApi(api models.TreatmentUpdateToUser) dto.TreatmentUpdateToUser {
	return dto.TreatmentUpdateToUser{
		ID:            api.ID,
		DoctorID:      api.DoctorID,
		Weight:        api.Weight,
		Temperature:   api.Temperature,
		Comment:       api.Comment,
		Prescriptions: mapDtoPrescriptionForUpdateOfApi(api.Prescriptions),
	}
}

func mapDtoPrescriptionForUpdateOfApi(api []models.PrescriptionForUpdate) []dto.PrescriptionForUpdate {
	prescriptionDTO := make([]dto.PrescriptionForUpdate, len(api))

	for i, p := range api {
		prescriptionDTO[i] = dto.PrescriptionForUpdate{
			PreparationID: p.PreparationID,
			Name:          p.Name,
			Dose:          p.Dose,
			Course:        p.Course,
			Category:      p.Category,
			Option:        p.Option,
		}
	}

	return prescriptionDTO
}
