package private

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// PreparationsToSymptomsHandler Препараты по симптомам
//
//	@Summary		Получить список препаратов по симптомам
//	@Description	Получить список препаратов по симптомам
//	@ID				get_preparations_to_symptoms
//	@Tags			preparations
//	@Accept			json
//	@Produce		json
//	@Param			ids	query		array							true	"список id симптомов"
//	@Success		200	{array}		models.PreparationsToSymptoms	"Список препаратов по симптомам"
//	@Failure		400	{object}	models.Response					"Ошибка запроса"
//	@Failure		500	{object}	models.Response					"Внутренняя ошибка сервера"
//	@Router			/preparationsToSymptoms [get]
func (s *Server) PreparationsToSymptomsHandler(c *fiber.Ctx) error {
	idsStr := c.Query("ids")
	idStrings := strings.Split(idsStr, ",")

	ids := make([]int64, len(idStrings))
	for _, idStr := range idStrings {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:        fiber.StatusBadRequest,
				Message:     "Invalid ID format",
				Description: "IDs must be integers separated by commas",
			})
		}
		ids = append(ids, id)
	}

	preparationsData, err := s.svc.GetPreparationsToSymptoms(c.Context(), ids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	preparations := make([]models.PreparationsToSymptoms, len(preparationsData))
	for i, p := range preparationsData {
		preparations[i] = models.PreparationsToSymptoms{
			ID:       p.ID,
			Name:     p.Name,
			Dose:     p.Dose,
			Course:   p.Course,
			Category: p.Category,
			Option:   p.Option,
			Similar:  mapSimilarDTOToApi(p.Similar),
		}
	}

	return c.JSON(preparations)
}

func mapSimilarDTOToApi(list []dto.NameResponse) []models.NameResponse {
	similar := make([]models.NameResponse, len(list))
	for i, p := range list {
		similar[i] = models.NameResponse{
			ID:   p.ID,
			Name: p.Name,
		}
	}
	return similar
}
