package application

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/api/rest"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/config/logging"
	"github.com/lallison21/library_rest_service/internal/services"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	cfg     *config.Config
	logging *slog.Logger
	service *services.Service
}

func New() *Application {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	service := services.New()
	logger := logging.New(cfg.Logging)

	return &Application{
		cfg:     &cfg,
		logging: logger,
		service: service,
	}
}

func (a *Application) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

		<-s
		cancel()
	}()

	server := rest.New(&a.cfg.Server, a.logging, a.service)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		a.logging.Info("shutting down server...")
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		a.logging.Info(fmt.Sprintf("exit reason: %s \n", err))
	}
}
