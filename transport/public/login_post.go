package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// LoginHandler Авторизация
//
//	@Summary		Авторизация
//	@Description	Авторизация
//	@ID				auth
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.LoginRequest		true	"Запрос"
//
//	@Success		200		{object}	models.LoginResponse	"Успешный ответ"
//	@Failure		400		{object}	models.Response			"Ошибка запроса"
//	@Failure		500		{object}	models.Response			"Внутренняя ошибка сервера"
//	@Router			/login [post]
func (s *Server) LoginHandler(c *fiber.Ctx) error {
	var login models.LoginRequest
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	if err := validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Validation failed",
			Description: err.Error(),
		})
	}

	l, err := s.svc.Login(c.Context(), getDtoLoginOfApi(login))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Role:         l.Role,
		Name:         l.Name,
		UserID:       l.UserID,
		AccessToken:  l.AccessToken,
		RefreshToken: l.RefreshToken,
	})
}

func getDtoLoginOfApi(model models.LoginRequest) dto.LoginRequest {
	return dto.LoginRequest{
		Username: model.Username,
		Password: model.Password,
	}
}
