package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/api/rest/auth_handler"
	"github.com/lallison21/library_rest_service/internal/api/rest/status_handler"
	"github.com/lallison21/library_rest_service/internal/application"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/config/logging"
	"github.com/lallison21/library_rest_service/internal/repository/auth_repo"
	"github.com/lallison21/library_rest_service/internal/repository/status_repo"
	"github.com/lallison21/library_rest_service/internal/services/auth_service"
	"github.com/lallison21/library_rest_service/internal/services/status_service"
)

func main() {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
	log := logging.New(cfg.Logging)

	app := application.New(&cfg, log)

	statusRepo := status_repo.New()
	statusService := status_service.New(statusRepo)
	statusHandler := status_handler.New(log, statusService)
	app.Handlers.Status = statusHandler

	authRepo := auth_repo.New()
	authService := auth_service.New(authRepo)
	authHandler := auth_handler.New(authService, log)
	app.Handlers.Auth = authHandler

	app.RegisterHandlers()

	app.Run()
}
