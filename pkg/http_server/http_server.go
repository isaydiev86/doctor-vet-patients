package http_server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

//type Config struct {
//	Host         string        `yaml:"host"`
//	Port         int           `yaml:"port"`
//	IdleTimeout  time.Duration `yaml:"idle_timeout"`
//	ReadTimeout  time.Duration `yaml:"read_timeout"`
//	WriteTimeout time.Duration `yaml:"write_timeout"`
//}

type Server struct {
	server *fiber.App
	notify chan error
}

func New(app *fiber.App, address string) *Server {
	s := &Server{
		server: app,
		notify: make(chan error, 1),
	}

	go s.start(address)

	return s
}

func (s *Server) start(address string) {
	s.notify <- s.server.Listen(address)
	close(s.notify)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.ShutdownWithContext(ctx)
	if err != nil {
		//log.Error().Err(err).Msg("server - Close - s.server.Shutdown")
	}
}
