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

	app.Get("/swagger/*", swagger.HandlerDefault)
	/// TODO разделить на админ и общие

	grp := app.Group("/api/v1")

	grp.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
	grp.Put("/patient", func(c *fiber.Ctx) error {
		return PatientUpdateHandler(c, svc)
	})

	grp.Get("/treatments", func(c *fiber.Ctx) error {
		return TreatmentsHandler(c, svc)
	})
	grp.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return TreatmentHandler(c, svc)
	})

	grp.Get("/reference", func(c *fiber.Ctx) error {
		return ReferenceHandler(c, svc)
	})

	grp.Post("/login", func(c *fiber.Ctx) error {
		return LoginHandler(c, svc)
	})

	grp.Get("/users", func(c *fiber.Ctx) error {
		return UserHandler(c, svc)
	})

}
