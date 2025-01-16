package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/version"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Handlers struct {
	Status StatusHandler
	Auth   AuthHandler
}

type Application struct {
	cfg *config.Config
	log *slog.Logger

	ctx    context.Context
	cancel context.CancelFunc

	server *http.Server
	router *gin.Engine

	Handlers Handlers
}

func New(cfg *config.Config, log *slog.Logger) *Application {
	router := gin.New()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := &http.Server{}

	app := &Application{
		cfg:    cfg,
		log:    log,
		ctx:    ctx,
		cancel: cancel,
		server: server,
		router: router,
	}
	return app
}

func (a *Application) RegisterHandlers() {
	a.router.GET("/ping", a.Handlers.Status.Ping())

	authRoute := a.router.Group("/auth")
	authRoute.POST("/register", a.Handlers.Auth.Register())
	authRoute.POST("/login", a.Handlers.Auth.Login())
}

func (a *Application) Run() {
	go a.gracefulShutdown()

	addr := fmt.Sprintf("%s:%s", a.cfg.Server.Host, a.cfg.Server.Port)
	gin.SetMode(a.cfg.Server.GinMode)

	a.server.Addr = addr
	a.server.Handler = a.router

	a.log.Info("application started",
		"address", addr,
		"name", version.Name,
		"version", version.Version,
		"build_time", version.BuildTime,
	)
	if err := a.server.ListenAndServe(); err != nil {
		a.log.Error("application failed to start", err.Error())
	}
}

func (a *Application) gracefulShutdown() {
	signalCtx, stop := signal.NotifyContext(a.ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-signalCtx.Done()
	a.log.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(shutdownCtx); err != nil {
		a.log.Error("failed to gracefully shutdown server", "error", err.Error())
	}
}
