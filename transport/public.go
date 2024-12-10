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

	grp := app.Group("/api/v1/public")

	grp.Post("/login", func(c *fiber.Ctx) error {
		return LoginHandler(c, svc)
	})
}
