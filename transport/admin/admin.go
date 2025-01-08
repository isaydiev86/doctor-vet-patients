package admin

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/transport"
	"github.com/isaydiev86/doctor-vet-patients/transport/middlewares"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
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
	allowedRoles := []string{"admin"}

	admin := s.App.Group("/api/v1/admin",
		keycloak.TokenValidationMiddleware(s.keycloak, s.log),
		keycloak.RoleValidationMiddleware(s.keycloak, s.log, allowedRoles...),
	)

	// Использование middleware для парсинга и валидации
	admin.Post("/symptoms", middlewares.ParseAndValidateMiddleware(&models.NameAdd{}), func(c *fiber.Ctx) error {
		return s.SymptomAddHandler(c)
	})

	admin.Post("/preparations", middlewares.ParseAndValidateMiddleware(&models.PreparationsAdd{}), func(c *fiber.Ctx) error {
		return s.PreparationAddHandler(c)
	})

	admin.Post("/relationSymptomWithPreparation", middlewares.ParseAndValidateMiddleware(&models.RelationSymptomWithPreparation{}), func(c *fiber.Ctx) error {
		return s.RelationSymptomWithPreparationHandler(c)
	})

	admin.Post("/patient", middlewares.ParseAndValidateMiddleware(&models.Patient{}), func(c *fiber.Ctx) error {
		return s.PatientAddHandler(c)
	})

	admin.Put("/patient", middlewares.ParseAndValidateMiddleware(&models.Patient{}), func(c *fiber.Ctx) error {
		return s.PatientUpdateHandler(c)
	})

	admin.Get("/treatments", func(c *fiber.Ctx) error {
		return s.TreatmentsHandler(c)
	})

	admin.Get("/treatment/:id", func(c *fiber.Ctx) error {
		return s.TreatmentHandler(c)
	})

	admin.Put("/treatment", middlewares.ParseAndValidateMiddleware(&models.TreatmentSendForUser{}), func(c *fiber.Ctx) error {
		return s.TreatmentSendOnUserHandler(c)
	})

	admin.Get("/reference", func(c *fiber.Ctx) error {
		return s.ReferenceHandler(c)
	})

	admin.Get("/users", func(c *fiber.Ctx) error {
		return s.UsersHandler(c)
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
