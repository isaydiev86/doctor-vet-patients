package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

var validate = validator.New()

// ParseAndValidateMiddleware Middleware для парсинга и валидации тела запроса
func ParseAndValidateMiddleware(request interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:        fiber.StatusBadRequest,
				Message:     "Invalid request data",
				Description: "Failed to parse request body",
			})
		}

		if err := validate.Struct(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Code:        fiber.StatusBadRequest,
				Message:     "Validation failed",
				Description: err.Error(),
			})
		}

		// Сохраняем распаршенные данные в контекст
		c.Locals("parsedRequest", request)

		return c.Next()
	}
}
