package transport

import (
	"doctor-vet-patients/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App, svc service.IService) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Передаем сервис в обработчики
	app.Get("/patients", func(c *fiber.Ctx) error {
		return PatientsHandler(c, svc)
	})

	app.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
}
