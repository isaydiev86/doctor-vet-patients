package transport

import (
	"context"

	"doctor-vet-patients/internal/dto"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport/models"
	"github.com/gofiber/fiber/v2"
)

// TreatmentsHandler Получить список лечений
//
//	@Summary		Получить список лечений
//	@Description	Возвращает список всех лечений
//	@ID				get_treatments
//	@Tags			patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Patient	"Список лечений"
//	@Failure		400	{object}	models.Response	"Ошибка запроса"
//	@Failure		500	{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/treatments [get]
func TreatmentsHandler(c *fiber.Ctx, svc service.Service) error {
	ctx := context.Background()

	/// TODO учесть фильтры (по фио, по кличке(name), по статусу, лимит и офсет, дате)

	treatmentsData, err := svc.GetTreatments(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     "Error",
			Description: "Error for backend",
		})
	}

	treatments := make([]models.Treatment, len(treatmentsData))
	for i, p := range treatmentsData {
		treatments[i] = models.Treatment{
			ID:            p.ID,
			PatientID:     p.PatientID,
			DoctorID:      p.DoctorID,
			Status:        p.Status,
			CreatedAt:     p.CreatedAt,
			UpdatedAt:     p.UpdatedAt,
			BeginAt:       p.BeginAt,
			EndAt:         p.EndAt,
			Comment:       p.Comment,
			IsActive:      p.IsActive,
			Age:           p.Age,
			Weight:        p.Weight,
			Temperature:   p.Temperature,
			Patient:       getPatientOfDTO(p.Patient),
			Prescriptions: getPrescriptionOfDTO(p.Prescription),
		}
	}

	return c.JSON(treatments)
}

func getPatientOfDTO(dto dto.Patient) models.Patient {
	return models.Patient{
		ID:         dto.ID,
		DoctorID:   dto.DoctorID,
		Fio:        dto.Fio,
		Phone:      dto.Phone,
		Address:    dto.Address,
		Animal:     dto.Animal,
		Name:       dto.Name,
		Breed:      dto.Breed,
		Age:        dto.Age,
		Gender:     dto.Gender,
		Status:     dto.Status,
		IsNeutered: dto.IsNeutered,
	}
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
