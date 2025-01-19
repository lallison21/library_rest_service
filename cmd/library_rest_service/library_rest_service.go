package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/api/rest/handler"
	"github.com/lallison21/library_rest_service/internal/application"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/config/datebase/postgres"
	"github.com/lallison21/library_rest_service/internal/config/logging"
	"github.com/lallison21/library_rest_service/internal/repository"
	"github.com/lallison21/library_rest_service/internal/services"
	"github.com/lallison21/library_rest_service/internal/utils"
)

func main() {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	log := logging.New(cfg.Logging)
	pgConn := postgres.New(cfg.Postgres, log)

	passUtils := utils.NewPassword(&cfg.Password)

	app := application.New(&cfg, log)

	statusRepo := repository.NewStatus(pgConn)
	statusService := services.NewStatus(statusRepo)
	statusHandler := handler.NewStatus(log, statusService)
	app.Handlers.Status = statusHandler

	authRepo := repository.NewAuth(pgConn)
	authService := services.NewAuth(authRepo, passUtils)
	authHandler := handler.NewAuth(authService, log)
	app.Handlers.Auth = authHandler

	app.RegisterHandlers()

	app.Run()
}
