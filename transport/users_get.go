package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// UserHandler Получить список пользователей
//
//	@Summary		Получить список пользователей
//	@Description	Возвращает список пользователей
//	@ID				get_users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			role	query		string			false	"Фильтр по role"
//	@Success		200		{array}		models.User		"Список пользователей"
//	@Failure		400		{object}	models.Response	"Ошибка запроса"
//	@Failure		500		{object}	models.Response	"Внутренняя ошибка сервера"
//	@Router			/users [get]
func UserHandler(c *fiber.Ctx, svc service.Service) error {
	filters := dto.UserFilters{
		Role: c.Query("role"),
	}
	usersData, err := svc.GetUsers(c.Context(), filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	users := make([]models.User, len(usersData))
	for i, p := range usersData {
		users[i] = models.User{
			ID:     p.ID,
			UserID: p.UserID,
			Fio:    p.Fio,
			Role:   p.Role,
		}
	}

	return c.JSON(users)
}
