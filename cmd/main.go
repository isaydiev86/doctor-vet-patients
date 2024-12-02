package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"doctor-vet-patients/db"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/pkg/dbutil"
	"doctor-vet-patients/pkg/utils"
	"doctor-vet-patients/transport"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB  *dbutil.Config          `yaml:"db"`
	Srv *transport.ServerConfig `yaml:"server"`
}

func main() {
	ctx := context.Background()

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("Не удалось инициализировать конфигурацию: %v", err)
	}

	svr, err := run(ctx, cfg)
	if err != nil {
		log.Fatalf("Не удалось запустить приложение: %v", err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  cfg.Srv.IdleTimeout,
		ReadTimeout:  cfg.Srv.ReadTimeout,
		WriteTimeout: cfg.Srv.WriteTimeout,
	})
	transport.RegisterPublicRoutes(app, utils.FromPtr(svr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%d", cfg.Srv.Host, cfg.Srv.Port)
	go func() {
		if err := app.Listen(address); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Printf("Server is running on: %s", address)

	<-quit
	log.Println("Shutting down server...")

	ctxt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctxt); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
}

func run(ctx context.Context, cfg *Config) (*service.Service, error) {
	logger, err := initLogger()
	if err != nil {
		log.Printf("Не удалось инициализировать логгер: %v\n", err)
		return nil, err
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
		return nil, err
	}

	svc := service.New(service.Relation{DB: bd}, logger)

	//initRouterPublic(ctx, svc)

	logger.Info("Приложение запустилось!")

	return svc, nil
}

func initRouterPublic(ctxt context.Context, svc *service.Service) {
	/// TODO как сделать Server gracefully stopped общий на приложение
	//app := fiber.New()
	//transport.RegisterRoutes(app, *svc)
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//
	//go func() {
	//	if err := app.Listen(":3000"); err != nil {
	//		log.Fatalf("Error starting server: %v", err)
	//	}
	//}()
	//
	//log.Println("Server is running on http://localhost:3000")
	//
	//<-quit
	//log.Println("Shutting down server...")
	//
	//ctx, cancel := context.WithTimeout(ctxt, 5*time.Second)
	//defer cancel()
	//
	//if err := app.ShutdownWithContext(ctx); err != nil {
	//	log.Fatalf("Error shutting down server: %v", err)
	//}

	log.Println("Server gracefully stopped")
}

func initLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction() // или zap.NewDevelopment() для разработки
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации логгера: %w", err)
	}
	return logger, nil
}

func initConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	return &config, nil
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
