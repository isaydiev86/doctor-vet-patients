package keycloak

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func TokenValidationMiddleware(k *Service, logger Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := getToken(c)
		if err != nil {
			logger.Warn("Authorization error: ", err)
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:        http.StatusUnauthorized,
				Message:     "Unauthorized",
				Description: err.Error(),
			})
		}

		valid, err := k.ValidateToken(c.Context(), token)
		if err != nil || !valid {
			logger.Warn("Invalid or expired token", err)
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
		token, err := getToken(c)
		if err != nil {
			logger.Warn("Authorization error: ", err)
			return c.Status(fiber.StatusUnauthorized).JSON(Response{
				Code:        http.StatusUnauthorized,
				Message:     "Unauthorized",
				Description: err.Error(),
			})
		}

		roles, err := k.GetUserRoles(token)
		if err != nil {
			logger.Warn("Failed to get user roles", err)
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

		logger.Warn("User does not have the required role", "required_role", allowedRoles)
		return c.Status(fiber.StatusForbidden).JSON(Response{
			Code:        http.StatusUnauthorized,
			Message:     "Forbidden: insufficient permissions",
			Description: "Forbidden: insufficient permissions",
		})
	}
}

func getToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	if token == "" {
		return "", errors.New("invalid Authorization header format")
	}

	return token, nil
}
