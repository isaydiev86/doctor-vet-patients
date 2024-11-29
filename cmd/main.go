package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"doctor-vet-patients/db"
	"doctor-vet-patients/internal/service"
	"doctor-vet-patients/transport"
)

type Config struct {
	DB db.Config
}

//TODO:
// take transaction away
// remove pointers

func main() {

	var cfg Config
	_ = cfg

	app := fiber.New()

	// Инициализация хранилища
	storage, err := db.New(cfg.DB)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cannot create application"))
	}
	//
	//// Инициализация сервиса
	svc := service.New(service.Relation{DB: storage})
	//
	//// Регистрация маршрутов с передачей сервиса
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
