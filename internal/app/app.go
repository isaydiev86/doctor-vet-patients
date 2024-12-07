package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/isaydiev86/doctor-vet-patients/db"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	"github.com/isaydiev86/doctor-vet-patients/pkg/dbutil"
	"github.com/isaydiev86/doctor-vet-patients/pkg/http_server"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/isaydiev86/doctor-vet-patients/transport"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	DB       *dbutil.Config    `yaml:"db"`
	Srv      *transport.Config `yaml:"server"`
	Keycloak *keycloak.Config  `yaml:"keycloak"`
}

func Run(cfg *Config, logger *zap.Logger) error {
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

	s := http_server.New(app, fmt.Sprintf("%s:%d", cfg.Srv.Host, cfg.Srv.Port))
	defer s.Close()

	waiting(logger)

	return nil
}

func waiting(logger *zap.Logger) {
	logger.Info("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	<-wait

	logger.Info("App is stopping...")
}

func initDB(ctx context.Context, cfg *Config, logger *zap.Logger) (*db.DB, error) {
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
