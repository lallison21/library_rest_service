package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/api/rest/auth_handler"
	"github.com/lallison21/library_rest_service/internal/api/rest/status_handler"
	"github.com/lallison21/library_rest_service/internal/application"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/config/datebase/postgres"
	"github.com/lallison21/library_rest_service/internal/config/logging"
	"github.com/lallison21/library_rest_service/internal/repository/auth_repo"
	"github.com/lallison21/library_rest_service/internal/repository/status_repo"
	"github.com/lallison21/library_rest_service/internal/services/auth_service"
	"github.com/lallison21/library_rest_service/internal/services/status_service"
	"github.com/lallison21/library_rest_service/internal/utils/password_utils"
)

func main() {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
	log := logging.New(cfg.Logging)
	pg := postgres.New(cfg.Postgres, log)

	passUtils := password_utils.New(&cfg.Password)

	app := application.New(&cfg, log)

	statusRepo := status_repo.New(pg)
	statusService := status_service.New(statusRepo)
	statusHandler := status_handler.New(log, statusService)
	app.Handlers.Status = statusHandler

	authRepo := auth_repo.New(pg)
	authService := auth_service.New(authRepo, passUtils)
	authHandler := auth_handler.New(authService, log)
	app.Handlers.Auth = authHandler

	app.RegisterHandlers()

	app.Run()
}
