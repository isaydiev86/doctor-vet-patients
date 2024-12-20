package public

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/isaydiev86/doctor-vet-patients/transport"
)

// RegisterPublicRoutes @title			Patient Service API
// @version		1.0
// @description	API для работы с пациентами и их данными
// @schemes		http
// @termsOfService	http://swagger.io.terms/

func New(cfg transport.Config, svc Services, log Logger) (*Server, error) {
	s := Server{
		log: log,
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

func (s *Server) Start(_ context.Context) error {
	s.App.Get("/swagger/*", swagger.HandlerDefault)

	grp := s.App.Group("/api/v1/public")

	grp.Post("/login", func(c *fiber.Ctx) error {
		return s.LoginHandler(c)
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
