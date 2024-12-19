package public

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/isaydiev86/doctor-vet-patients/pkg/logger"
	"github.com/isaydiev86/doctor-vet-patients/transport"
)

// RegisterPublicRoutes @title			Patient Service API
// @version		1.0
// @description	API для работы с пациентами и их данными
// @schemes		http
// @termsOfService	http://swagger.io.terms/

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
	s.App.Get("/swagger/*", swagger.HandlerDefault)

	grp := s.App.Group("/api/v1/public")

	grp.Post("/login", func(c *fiber.Ctx) error {
		return s.LoginHandler(c)
	})

	return s.App.Listen(s.cfg.Host + ":" + string(s.cfg.Port))
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Sync()
	return s.App.Shutdown()
}
