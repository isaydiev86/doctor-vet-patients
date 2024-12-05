package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/isaydiev86/doctor-vet-patients/internal/app"
	"gopkg.in/yaml.v3"
)

func main() {
	ctx := context.Background()

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("Не удалось инициализировать конфигурацию: %v", err)
	}

	err = app.Run(ctx, cfg)
	if err != nil {
		log.Fatalf("Не удалось запустить приложение: %v", err)
	}

	/// TODO как сделать Server gracefully stopped общий на приложение

	//app := fiber.New(fiber.Config{
	//	IdleTimeout:  cfg.Srv.IdleTimeout,
	//	ReadTimeout:  cfg.Srv.ReadTimeout,
	//	WriteTimeout: cfg.Srv.WriteTimeout,
	//})
	//transport.RegisterPublicRoutes(app, utils.FromPtr(svr))
	//
	//
	//
	//address := fmt.Sprintf("%s:%d", cfg.Srv.Host, cfg.Srv.Port)
	//go func() {
	//	if err := app.Listen(address); err != nil {
	//		log.Fatalf("Error starting server: %v", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//
	//log.Printf("Server is running on: %s", address)
	//
	//<-quit
	//log.Println("Shutting down server...")
	//
	//ctxt, cancel := context.WithTimeout(ctx, 5*time.Second)
	//defer cancel()
	//
	//if err := app.ShutdownWithContext(ctxt); err != nil {
	//	log.Fatalf("Error shutting down server: %v", err)
	//}
	//
	//log.Println("Server gracefully stopped")
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
