package main

import (
	"fmt"
	"log"

	"github.com/isaydiev86/doctor-vet-patients/config"
	"github.com/isaydiev86/doctor-vet-patients/internal/app"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать конфигурацию: %v", err)
	}

	logger, err := initLogger()
	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v\n", err)
	}

	err = app.Run(cfg, logger)
	if err != nil {
		log.Fatalf("Не удалось запустить приложение: %v", err)
	}
}

func initLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction() // или zap.NewDevelopment() для разработки
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации логгера: %w", err)
	}
	return logger, nil
}
