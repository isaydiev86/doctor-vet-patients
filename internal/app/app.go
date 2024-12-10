package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/config"
	"github.com/isaydiev86/doctor-vet-patients/db"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/pkg/http_server"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/isaydiev86/doctor-vet-patients/transport"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	ctx := context.Background()

	bd, err := initDB(ctx, cfg, logger)
	if err != nil {
		logger.Error("cannot create application", zap.Error(err))
		return err
	}
	defer func(bd *db.DB, ctx context.Context) {
		err = bd.Stop(ctx)
		if err != nil {
			logger.Error("cannot close db", zap.Error(err))
		}
	}(bd, ctx)

	kcConfig := keycloak.Config{
		URL:      cfg.Keycloak.URL,
		Realm:    cfg.Keycloak.Realm,
		ClientID: cfg.Keycloak.ClientID,
		Secret:   cfg.Keycloak.Secret,
	}
	keycloakService := keycloak.New(kcConfig)

	svc := service.New(service.Relation{DB: bd}, logger, keycloakService)

	app := fiber.New(fiber.Config{
		IdleTimeout:  cfg.Srv.IdleTimeout,
		ReadTimeout:  cfg.Srv.ReadTimeout,
		WriteTimeout: cfg.Srv.WriteTimeout,
	})

	transport.RegisterPublicRoutes(app, utils.FromPtr(svc))

	transport.RegisterAdminRoutes(app, utils.FromPtr(svc))

	transport.RegisterPrivateRoutes(app, utils.FromPtr(svc), true)

	s := http_server.New(app, fmt.Sprintf("%s:%d", cfg.Srv.Host, cfg.Srv.Port))
	defer s.Close()

	waiting(logger)

	return nil
}

//func (s *Server) Start(ctx context.Context) error {
//	timeoutMw, err := fiberutil.TimeoutMiddleware(fiberutil.WithTimeoutConfig(s.cfg.Timeout))
//	if err != nil {
//		return errors.Wrap(err, "new timeout middleware")
//	}
//	s.Use(timeoutMw)
//	s.Use(fiberutil.RequestIDMiddleware())
//	s.Use(s.Wrap())
//	metricsMiddleware, err := fiberutil.PrometheusMetricsMiddleware("", fiberutil.DefaultSubsystem)
//	if err != nil {
//		return errors.Wrap(err, "new metrics middleware")
//	}
//	s.Use(metricsMiddleware)
//	s.Get("/v1/", s.getAllProducts)
//	s.Post("/v1/calc-fee", s.postCalcFee)
//	return s.Server.Start(ctx)
//}

func waiting(logger *zap.Logger) {
	logger.Info("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	<-wait

	logger.Info("App is stopping...")
}

func initDB(ctx context.Context, cfg *config.Config, logger *zap.Logger) (*db.DB, error) {
	bd, err := db.New(utils.FromPtr(cfg.DB), logger)
	if err != nil {
		logger.Error("Ошибка создания базы данных", zap.Error(err))
		return nil, errors.Wrap(err, "cannot create db")
	}

	err = bd.DB.Start(ctx)
	if err != nil {
		logger.Error("Ошибка запуска базы данных", zap.Error(err))
		return nil, err
	}
	logger.Info("База данных успешно инициализирована")

	return bd, err
}
