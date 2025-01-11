package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

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
	treatment, ok := c.Locals("parsedRequest").(*models.TreatmentUpdateToUser)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Internal server error",
			Description: "Failed to parse request data",
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

func mapDtoUpdateTreatmentOfApi(api *models.TreatmentUpdateToUser) dto.TreatmentUpdateToUser {
	return dto.TreatmentUpdateToUser{
		ID:            api.ID,
		DoctorID:      api.DoctorID,
		Weight:        api.Weight,
		Temperature:   api.Temperature,
		Comment:       api.Comment,
		Prescriptions: mapDtoPrescriptionForUpdateOfApi(api.Prescriptions),
		AddInfo:       mapAddInfoOfApi(api.AddInfo),
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

func mapAddInfoOfApi(api []models.AddInfo) []dto.AddInfo {
	addInfo := make([]dto.AddInfo, len(api))
	for i, a := range api {
		addInfo[i] = dto.AddInfo{
			Key:      a.Key,
			Value:    a.Value,
			DataType: a.DataType,
			Name:     a.Name,
		}
	}
	return addInfo
}
