package app

import (
	"context"
	"fmt"
	"log"

	"doctor-vet-patients/db"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/pkg/dbutil"
	"doctor-vet-patients/pkg/utils"
	"doctor-vet-patients/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	DB  *dbutil.Config          `yaml:"db"`
	Srv *transport.ServerConfig `yaml:"server"`
}

func Run(ctx context.Context, cfg *Config) error {
	logger, err := initLogger()
	if err != nil {
		log.Printf("Не удалось инициализировать логгер: %v\n", err)
		return err
	}
	//defer func(logger *zap.Logger) {
	//	err := logger.Sync()
	//	if err != nil {
	//		log.Fatalf("Не удалось закрыть логгер: %v", err)
	//	}
	//}(logger)
	//
	//undo := zap.RedirectStdLog(logger)
	//defer undo()

	bd, err := initDB(ctx, cfg, logger)
	if err != nil {
		logger.Error("cannot create application", zap.Error(err))
		return err
	}

	svc := service.New(service.Relation{DB: bd}, logger)

	err = initRouterPublic(cfg, svc)
	if err != nil {
		logger.Error("cannot init router", zap.Error(err))
		return err
	}

	logger.Info("Приложение запустилось!")

	return nil
}

func initRouterPublic(cfg *Config, svc *service.Service) error {

	app := fiber.New(fiber.Config{
		IdleTimeout:  cfg.Srv.IdleTimeout,
		ReadTimeout:  cfg.Srv.ReadTimeout,
		WriteTimeout: cfg.Srv.WriteTimeout,
	})

	transport.RegisterPublicRoutes(app, utils.FromPtr(svc))

	address := fmt.Sprintf("%s:%d", cfg.Srv.Host, cfg.Srv.Port)

	return app.Listen(address)
}

func initLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction() // или zap.NewDevelopment() для разработки
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации логгера: %w", err)
	}
	return logger, nil
}

func initDB(ctx context.Context, cfg *Config, logger *zap.Logger) (*db.DB, error) {
	bd, err := db.New(utils.FromPtr(cfg.DB), logger)
	if err != nil {
		logger.Error("Ошибка создания базы данных", zap.Error(err))
		return nil, errors.Wrap(err, "cannot create application")
	}

	err = bd.DB.Start(ctx)
	if err != nil {
		logger.Error("Ошибка запуска базы данных", zap.Error(err))
		return nil, err
	}
	logger.Info("База данных успешно инициализирована")

	return bd, err
}
