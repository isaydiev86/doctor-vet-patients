package private

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/pkg/logger"
	"github.com/isaydiev86/doctor-vet-patients/transport"
)

func New(cfg transport.Config, svc Services) (*Server, error) {
	logger, err := logger.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v\n", err)
	}
	s := Server{
		log: logger,
		svc: svc,
		cfg: cfg,
	}
	s.App = fiber.New(fiber.Config{
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	return &s, nil
}

type Server struct {
	*fiber.App
	log Logger
	svc Services
	cfg transport.Config
}

func (s *Server) Start(ctx context.Context) error {
	/*allowedRoles := []string{"doctor"}
	if isAdmin {
		allowedRoles = append(allowedRoles, "admin")
	}
	/**/
	allowedRoles := []string{"doctor"}
	allowedRoles = append(allowedRoles, "admin")

	private := s.App.Group("/api/v1/private",
		keycloak.TokenValidationMiddleware(s.Keycloak, s.log),
		keycloak.RoleValidationMiddleware(svc.Keycloak, s.log, allowedRoles...),
	)

	private.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return s.TreatmentHandler(c)
	})

	private.Get("/reference", func(c *fiber.Ctx) error {
		return s.ReferenceHandler(c)
	})

	return s.App.Listen(s.cfg.Host + ":" + string(s.cfg.Port))
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Sync()
	return s.App.Shutdown()
}
