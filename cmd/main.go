package main

import (
	"log"
	_ "subs-service/docs"
	apiHTTP "subs-service/internal/api/http"
	"subs-service/internal/config"
	repo "subs-service/internal/repository/postgres"
	"subs-service/internal/usecases/service"
	pkgConfig "subs-service/pkg/config"
	"subs-service/pkg/database/postgres"
	"subs-service/pkg/http/handlers"
	"subs-service/pkg/http/server"

	"github.com/go-chi/chi/v5"
)

// @title Subscriptions Service API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1
func main() {
	appFlags := pkgConfig.ParseFlags()
	var cfg config.Config
	pkgConfig.MustLoadConfig(appFlags.ConfigPath, &cfg)

	log.Printf("subs-service server is starting")

	pool, err := postgres.NewPostgresPool(cfg.PostgresCfg)
	if err != nil {
		log.Fatalf("Failed to connect PostgreSQL: %s", err.Error())
	}

	log.Printf("Connected to PostgreSQL successfully")

	subsRepo := repo.NewSubsRepo(pool)
	subService := service.NewSubService(subsRepo)
	subHandler := apiHTTP.NewSubHandler(subService, cfg.PathCfg, cfg.SvcCfg, cfg.DataCfg)

	log.Printf("All services were created successfully")

	r := chi.NewRouter()
	handlers.RouteHandlers(r, cfg.PathCfg.API,
		handlers.WithLogger(),
		handlers.WithRecovery(),
		handlers.WithSwagger(),
		handlers.WithHealthHandler(),
		subHandler.WithSubHandlers(),
	)

	log.Printf("Starting HTTP server at %s...", cfg.HTTPCfg.Address)

	if err = server.CreateServer(r, cfg.HTTPCfg); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
