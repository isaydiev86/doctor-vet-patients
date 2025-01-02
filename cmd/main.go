package main

import (
	"context"
	"log"

	"github.com/isaydiev86/doctor-vet-patients/config"
	"github.com/isaydiev86/doctor-vet-patients/internal/service"
	adminRout "github.com/isaydiev86/doctor-vet-patients/transport/admin"
	privateRout "github.com/isaydiev86/doctor-vet-patients/transport/private"
	publicRout "github.com/isaydiev86/doctor-vet-patients/transport/public"

	"github.com/isaydiev86/doctor-vet-patients/pkg/app"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	zapLogger "github.com/isaydiev86/doctor-vet-patients/pkg/logger"

	"github.com/isaydiev86/doctor-vet-patients/db"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error init config: %v", err)
	}

	logger, err := zapLogger.New()
	if err != nil {
		logger.Fatal("Error init logger: %v\n", err)
	}

	bd, err := db.New(cfg.DB, logger)
	if err != nil {
		logger.Fatal("Error create database", err)
	}
	kcConfig := keycloak.Config{
		URL:      cfg.Keycloak.URL,
		Realm:    cfg.Keycloak.Realm,
		ClientID: cfg.Keycloak.ClientID,
		Secret:   cfg.Keycloak.Secret,
	}
	keycloakService := keycloak.New(kcConfig)

	svc := service.New(service.Relation{DB: bd}, logger, keycloakService)

	public := publicRout.New(cfg.Public, svc, logger)

	private := privateRout.New(cfg.Private, svc, logger, keycloakService)

	admin := adminRout.New(cfg.Admin, svc, logger, keycloakService)

	theApp, err := app.New(
		logger,
		app.NewLifecycleComponent("db", bd),
		app.NewLifecycleComponent("public", public),
		app.NewLifecycleComponent("private", private),
		app.NewLifecycleComponent("admin", admin),
	)
	if err != nil {
		logger.Fatal("Error create application", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = theApp.Run(ctx)
	if err != nil {
		logger.Fatal("Error run app: %v", err)
	}
}
