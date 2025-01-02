package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// TreatmentForUserHandler Получить детали лечения для доктора на кого назначено
//
//	@Summary		Получить детали лечения для доктора
//	@Description	Возвращает детали  лечения для доктора
//	@ID				get_treatment_for_user
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.TreatmentDetail	"Детали лечения"
//	@Failure		400	{object}	models.Response			"Ошибка запроса"
//	@Failure		500	{object}	models.Response			"Внутренняя ошибка сервера"
//	@Router			/treatment [get]
func (s *Server) TreatmentForUserHandler(c *fiber.Ctx) error {

	token, err := s.keycloak.GetToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Code:        fiber.StatusUnauthorized,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	userID, err := s.keycloak.ExtractUserIDFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Code:        fiber.StatusUnauthorized,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	treatmentsDto, err := s.svc.GetTreatmentForUser(c.Context(), userID)
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
		Weight:        treatmentsDto.Weight,
		Temperature:   treatmentsDto.Temperature,
		Patient:       getPatientOfDTO(treatmentsDto.Patient),
		Prescriptions: getPrescriptionOfDTO(treatmentsDto.Prescription),
	}

	return c.JSON(treatment)
}
