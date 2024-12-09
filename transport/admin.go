package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
)

func RegisterAdminRoutes(app *fiber.App, svc service.Service) {
	//tokenRetrospector := keycloak.New(utils.FromPtr(cfg.Keycloak))
	//app.Use(middlewares.NewJwtMiddleware(tokenRetrospector, utils.FromPtr(cfg.Keycloak)))

	grp := app.Group("/api/v1/admin")

	grp.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
	grp.Put("/patient", func(c *fiber.Ctx) error {
		return PatientUpdateHandler(c, svc)
	})

	grp.Get("/treatments", func(c *fiber.Ctx) error {
		return TreatmentsHandler(c, svc)
	})

	grp.Get("/users", func(c *fiber.Ctx) error {
		return UserHandler(c, svc)
	})

}
