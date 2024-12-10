package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
)

func RegisterPrivateRoutes(app *fiber.App, svc service.Service, isAdmin bool) {
	allowedRoles := []string{"doctor"}
	if isAdmin {
		allowedRoles = append(allowedRoles, "admin")
	}

	private := app.Group("/api/v1/private",
		keycloak.TokenValidationMiddleware(svc.Keycloak, svc.Logger),
		keycloak.RoleValidationMiddleware(svc.Keycloak, svc.Logger, allowedRoles...),
	)

	private.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return TreatmentHandler(c, svc)
	})

	private.Get("/reference", func(c *fiber.Ctx) error {
		return ReferenceHandler(c, svc)
	})

}
