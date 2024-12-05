package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
)

// RegisterPublicRoutes @title			Patient Service API
// @version		1.0
// @description	API для работы с пациентами и их данными
// @schemes		http
// @termsOfService	http://swagger.io.terms/
func RegisterPublicRoutes(app *fiber.App, svc service.Service) {

	/// TODO разделить на админ и общие

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
	app.Put("/patient", func(c *fiber.Ctx) error {
		return PatientUpdateHandler(c, svc)
	})

	app.Get("/treatments", func(c *fiber.Ctx) error {
		return TreatmentsHandler(c, svc)
	})
	app.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return TreatmentHandler(c, svc)
	})

	app.Get("/reference", func(c *fiber.Ctx) error {
		return ReferenceHandler(c, svc)
	})

	/// TODO пока не работает...
	app.Post("/login", func(c *fiber.Ctx) error {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&credentials); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
		}

		token, err := svc.Keycloak.Login(c.Context(), credentials.Username, credentials.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}

		return c.JSON(fiber.Map{"access_token": token})
	})

}
