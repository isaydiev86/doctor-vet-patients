package transport

import (
	"doctor-vet-patients/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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

}
