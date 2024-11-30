package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"doctor-vet-patients/pkg/dbutil"
	"doctor-vet-patients/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"doctor-vet-patients/db"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport"
)

type Config struct {
	DB *dbutil.Config `yaml:"db"`
}

func main() {
	ctx := context.Background()

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("Не удалось инициализировать конфигурацию: %v", err)
	}

	bd, err := initDB(ctx, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot create application"))
	}

	//// Инициализация сервиса
	svc := service.New(service.Relation{DB: bd})

	//// Регистрация маршрутов с передачей сервиса
	app := fiber.New()
	transport.RegisterRoutes(app, *svc)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Println("Server is running on http://localhost:3000")

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
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

func initDB(ctx context.Context, cfg *Config) (*db.DB, error) {
	// Инициализация хранилища
	bd, err := db.New(utils.FromPtr(cfg.DB))
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot create application"))
	}

	err = bd.DB.Start(ctx)
	if err != nil {
		return nil, err
	}

	return bd, err
}
