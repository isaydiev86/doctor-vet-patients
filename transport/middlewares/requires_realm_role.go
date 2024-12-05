package middlewares

import (
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
)

func NewRequiresRealmRole(role string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		claims := ctx.Value(keycloak.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := keycloak.NewJwtHelper(claims)
		if !jwtHelper.IsUserInRealmRole(role) {
			return c.Status(fiber.StatusUnauthorized).SendString("role authorization failed")
		}
		return c.Next()
	}
}
