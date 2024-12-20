package main

import (
	"context"
	"log"

	"github.com/isaydiev86/doctor-vet-patients/config"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	privateRout "github.com/isaydiev86/doctor-vet-patients/transport/private"
	publicRout "github.com/isaydiev86/doctor-vet-patients/transport/public"

	"github.com/isaydiev86/doctor-vet-patients/pkg/app"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/pkg/logger"

	"github.com/isaydiev86/doctor-vet-patients/db"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать конфигурацию: %v", err)
	}

	logger, err := logger.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v\n", err)
	}

	db, err := db.New(cfg.DB, logger)
	if err != nil {
		logger.Fatal("Ошибка создания базы данных", err)
	}
	kcConfig := keycloak.Config{
		URL:      cfg.Keycloak.URL,
		Realm:    cfg.Keycloak.Realm,
		ClientID: cfg.Keycloak.ClientID,
		Secret:   cfg.Keycloak.Secret,
	}
	keycloakService := keycloak.New(kcConfig)

	svc := service.New(service.Relation{DB: db}, logger, keycloakService)

	public, err := publicRout.New(cfg.Public, svc)
	if err != nil {
		logger.Fatal("Ошибка создания public", err)
	}

	private, err := privateRout.New(cfg.Private, svc, keycloakService)
	if err != nil {
		logger.Fatal("Ошибка создания private", err)
	}
	theApp, err := app.New(
		logger,
		app.NewLifecycleComponent("db", db),
		app.NewLifecycleComponent("public", public),
		app.NewLifecycleComponent("private", private),
	)
	if err != nil {
		logger.Fatal("Ошибка создания application", err)
	}

	ctx := context.Background()
	err = theApp.Run(ctx)
	if err != nil {
		log.Fatalf("Не удалось запустить приложение: %v", err)
	}
}
