package private

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/transport"
)

func New(cfg transport.Config, svc Services, log Logger, keycloak *keycloak.Service) *Server {
	s := Server{
		log:      log,
		svc:      svc,
		keycloak: keycloak,
		cfg:      cfg,
	}
	s.App = fiber.New(fiber.Config{
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	return &s
}

type Server struct {
	*fiber.App
	log      Logger
	svc      Services
	keycloak *keycloak.Service
	cfg      transport.Config
}

func (s *Server) Start(_ context.Context) error {
	allowedRoles := []string{"doctor", "admin"}

	private := s.App.Group("/api/v1/private",
		keycloak.TokenValidationMiddleware(s.keycloak, s.log),
		keycloak.RoleValidationMiddleware(s.keycloak, s.log, allowedRoles...),
	)

	private.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return s.TreatmentHandler(c)
	})

	private.Get("/treatment", func(c *fiber.Ctx) error {
		return s.TreatmentForUserHandler(c)
	})

	private.Put("/treatment", func(c *fiber.Ctx) error {
		return s.TreatmentUpdateHandler(c)
	})

	private.Get("/reference", func(c *fiber.Ctx) error {
		return s.ReferenceHandler(c)
	})

	private.Get("/symptoms", func(c *fiber.Ctx) error {
		return s.SymptomsHandler(c)
	})

	private.Get("/preparations", func(c *fiber.Ctx) error {
		return s.PreparationsHandler(c)
	})

	private.Get("/preparationsToSymptoms", func(c *fiber.Ctx) error {
		return s.PreparationsToSymptomsHandler(c)
	})

	return s.App.Listen(s.cfg.Host + ":" + strconv.Itoa(s.cfg.Port))
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.log.Sync()
	if err != nil {
		return err
	}
	return s.App.ShutdownWithContext(ctx)
}
