package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"doctor-vet-patients/db/postgres"
	"doctor-vet-patients/docs"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport"

	"github.com/gofiber/fiber/v2"
)

// @title			Patient Service API
// @version		1.0
// @description	API для работы с пациентами и их данными
// @schemes		http
// @termsOfService	http://swagger.io.terms/
func main() {
	app := fiber.New()

	docs.SwaggerInfo.BasePath = "/"

	// Инициализация хранилища
	storage := postgres.NewRepoPostgres()

	// Инициализация сервиса
	svc := service.NewService(storage)

	// Регистрация маршрутов с передачей сервиса
	transport.RegisterRoutes(app, svc)

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
