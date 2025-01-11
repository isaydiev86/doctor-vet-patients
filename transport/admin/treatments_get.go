package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// TreatmentsHandler Получить список лечений
//
//	@Summary		Получить список лечений
//	@Description	Возвращает список всех лечений
//	@ID				get_treatments
//	@Tags			treatment
//	@Accept			json
//	@Produce		json
//	@Param			fio		query		string				false	"Фильтр по ФИО"
//	@Param			name	query		string				false	"Фильтр по кличке"
//	@Param			status	query		string				false	"Фильтр по статусу"
//	@Param			date	query		string				false	"Фильтр по дате создания"
//	@Param			limit	query		integer				false	"Лимит записей (по умолчанию 10)"
//	@Param			offset	query		integer				false	"Смещение для пагинации (по умолчанию 0)"
//	@Success		200		{array}		models.Treatment	"Список лечений"
//	@Failure		400		{object}	models.Response		"Ошибка запроса"
//	@Failure		500		{object}	models.Response		"Внутренняя ошибка сервера"
//	@Router			/treatments [get]
func (s *Server) TreatmentsHandler(c *fiber.Ctx) error {
	filters := dto.TreatmentFilters{
		Fio:    c.Query("fio"),
		Name:   c.Query("name"),
		Status: c.Query("status"),
		Date:   c.Query("date"),
		Limit:  c.QueryInt("limit", 10),
		Offset: c.QueryInt("offset", 0),
	}
	treatmentsData, err := s.svc.GetTreatments(c.Context(), filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	treatments := make([]models.Treatment, len(treatmentsData))
	for i, p := range treatmentsData {
		treatments[i] = models.Treatment{
			ID:          p.ID,
			PatientID:   p.PatientID,
			DoctorID:    p.DoctorID,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			BeginAt:     p.BeginAt,
			EndAt:       p.EndAt,
			Comment:     p.Comment,
			IsActive:    p.IsActive,
			Weight:      p.Weight,
			Temperature: p.Temperature,
			Patient:     getPatientOfDTO(p.Patient),
		}
	}

	return c.JSON(treatments)
}

func getPatientOfDTO(dto dto.Patient) models.Patient {
	return models.Patient{
		ID:         dto.ID,
		Fio:        dto.Fio,
		Phone:      dto.Phone,
		Animal:     utils.ToPtr(dto.Animal),
		Name:       dto.Name,
		Breed:      utils.ToPtr(dto.Breed),
		Age:        utils.ToPtr(dto.Age),
		Gender:     dto.Gender,
		IsNeutered: dto.IsNeutered,
	}
}
