package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

// RefreshTokenHandler Обновление токена
//
//	@Summary		Обновление токена
//	@Description	Обновление токена
//	@ID				refreshToken
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
//	@Param			Form	body		models.RefreshTokenRequest	true	"Запрос"
//
//	@Success		200		{object}	models.RefreshTokenShort	"Успешный ответ"
//	@Failure		400		{object}	models.Response				"Ошибка запроса"
//	@Failure		500		{object}	models.Response				"Внутренняя ошибка сервера"
//	@Router			/refreshToken [post]
func (s *Server) RefreshTokenHandler(c *fiber.Ctx) error {
	var refresh models.RefreshTokenRequest
	if err := c.BodyParser(&refresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Invalid request data",
			Description: "Failed to parse request body",
		})
	}

	if err := validate.Struct(refresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Code:        fiber.StatusBadRequest,
			Message:     "Validation failed",
			Description: err.Error(),
		})
	}

	r, err := s.svc.RefreshToken(c.Context(), refresh.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			Description: err.Error(),
		})
	}

	return c.JSON(models.RefreshTokenShort{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	})
}
