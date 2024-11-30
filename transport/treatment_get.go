package transport

import (
	"context"
	"strconv"

	"doctor-vet-patients/internal/dto"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport/models"
	"github.com/gofiber/fiber/v2"
)

// TreatmentHandler Получить детали лечения
//
//	@Summary		Получить детали лечения
//	@Description	Возвращает детали  лечения
//	@ID				get_treatment
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.TreatmentDetail	"Детали лечения"
//	@Failure		400	{object}	models.Response	"Ошибка запроса"
//	@Failure		500	{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/treatment/{id} [get]
func TreatmentHandler(c *fiber.Ctx, svc service.Service) error {
	ctx := context.Background()

	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid ID parameter",
			Description: err.Error(),
		})
	}

	treatmentsDto, err := svc.GetTreatment(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	treatment := models.TreatmentDetail{
		ID:            treatmentsDto.ID,
		PatientID:     treatmentsDto.PatientID,
		DoctorID:      treatmentsDto.DoctorID,
		Status:        treatmentsDto.Status,
		CreatedAt:     treatmentsDto.CreatedAt,
		UpdatedAt:     treatmentsDto.UpdatedAt,
		BeginAt:       treatmentsDto.BeginAt,
		EndAt:         treatmentsDto.EndAt,
		Comment:       treatmentsDto.Comment,
		IsActive:      treatmentsDto.IsActive,
		Age:           treatmentsDto.Age,
		Weight:        treatmentsDto.Weight,
		Temperature:   treatmentsDto.Temperature,
		Patient:       getPatientOfDTO(treatmentsDto.Patient),
		Prescriptions: getPrescriptionOfDTO(treatmentsDto.Prescription),
	}

	return c.JSON(treatment)
}

func getPrescriptionOfDTO(dto []dto.Prescription) []models.Prescription {
	prescription := make([]models.Prescription, len(dto))
	for i, p := range dto {
		prescription[i] = models.Prescription{
			ID:          p.ID,
			TreatmentID: p.TreatmentID,
			Preparation: p.Preparation,
			Dose:        p.Dose,
			Course:      p.Course,
			Amount:      p.Amount,
		}
	}
	return prescription
}
