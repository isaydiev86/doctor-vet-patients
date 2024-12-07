package main

import (
	"fmt"
	"log"
	"os"

	"github.com/isaydiev86/doctor-vet-patients/internal/app"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func main() {

	cfg, err := initConfig()
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

func initConfig() (*app.Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var config app.Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	return &config, nil
}
