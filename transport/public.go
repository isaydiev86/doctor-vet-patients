package transport

import (
	"doctor-vet-patients/docs"
	"doctor-vet-patients/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// RegisterRoutes @title			Patient Service API
// @version		1.0
// @description	API для работы с пациентами и их данными
// @schemes		http
// @termsOfService	http://swagger.io.terms/
func RegisterRoutes(app *fiber.App, svc service.Service) {

	docs.SwaggerInfo.BasePath = "/"

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/patients", func(c *fiber.Ctx) error {
		return PatientsHandler(c, svc)
	})

	app.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
}
