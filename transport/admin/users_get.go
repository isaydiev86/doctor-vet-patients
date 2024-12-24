package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// UsersHandler Справочник
//
//	@Summary		Получить список пользователей
//	@Description	Получить список пользователей
//	@ID				get_users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			role	query		string				true	"роль"
//	@Success		200		{array}		models.User     	"Список пользователей"
//	@Failure		400		{object}	models.Response		"Ошибка запроса"
//	@Failure		500		{object}	models.Response		"Внутренняя ошибка сервера"
//	@Router			/users [get]
func (s *Server) UsersHandler(c *fiber.Ctx) error {
	role := c.Query("role", "doctor")

	userData, err := s.svc.GetUsers(c.Context(), role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	users := make([]models.User, len(userData))
	for i, p := range userData {
		users[i] = models.User{
			ID:     p.ID,
			UserID: p.UserID,
			Fio:    p.Fio,
			Role:   p.Role,
		}
	}

	return c.JSON(users)
}
