package keycloak

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Response struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func TokenValidationMiddleware(k *Service, logger Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			logger.Warn("Missing Authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(
				Response{
					Code:        http.StatusUnauthorized,
					Message:     "Missing Authorization header",
					Description: "Missing Authorization header",
				})
		}

		token := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		if token == "" {
			logger.Warn("Invalid Authorization header format")
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:        http.StatusUnauthorized,
				Message:     "Invalid Authorization header format",
				Description: "Invalid Authorization header format",
			})
		}

		valid, err := k.ValidateToken(c.Context(), token)
		if err != nil || !valid {
			logger.Warn("Invalid or expired token", zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:        http.StatusUnauthorized,
				Message:     "Invalid or expired token",
				Description: "Invalid or expired token",
			})
		}

		return c.Next()
	}
}

func RoleValidationMiddleware(k *Service, logger Logger, allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		token := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		roles, err := k.GetUserRoles(token)
		if err != nil {
			logger.Warn("Failed to get user roles", zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:        http.StatusUnauthorized,
				Message:     "Failed to get user roles",
				Description: "Failed to get user roles",
			})
		}

		for _, userRole := range roles {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					return c.Next()
				}
			}
		}

		logger.Warn("User does not have the required role", zap.Strings("required_role", allowedRoles))
		return c.Status(fiber.StatusForbidden).JSON(Response{
			Code:        http.StatusUnauthorized,
			Message:     "Forbidden: insufficient permissions",
			Description: "Forbidden: insufficient permissions",
		})
	}
}
