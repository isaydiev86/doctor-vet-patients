package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
)

func RegisterAdminRoutes(app *fiber.App, svc service.Service) {
	admin := app.Group("/api/v1/admin",
		keycloak.TokenValidationMiddleware(svc.Keycloak, svc.Logger),
		keycloak.RoleValidationMiddleware(svc.Keycloak, svc.Logger, "admin"),
	)

	admin.Post("/patient", func(c *fiber.Ctx) error {
		return PatientAddHandler(c, svc)
	})
	admin.Put("/patient", func(c *fiber.Ctx) error {
		return PatientUpdateHandler(c, svc)
	})

	admin.Get("/treatments", func(c *fiber.Ctx) error {
		return TreatmentsHandler(c, svc)
	})

	admin.Get("/users", func(c *fiber.Ctx) error {
		return UserHandler(c, svc)
	})
}
