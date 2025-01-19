package application

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/version"
)

type Handlers struct {
	Status StatusHandler
	Auth   AuthHandler
}

type Application struct {
	cfg *config.Config
	log *slog.Logger

	server *http.Server
	router *gin.Engine

	Handlers Handlers
}

func New(cfg *config.Config, log *slog.Logger) *Application {
	router := gin.New()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}

	server := &http.Server{
		ReadTimeout: cfg.Server.ReadHeaderTimeout,
	}

	app := &Application{
		cfg:    cfg,
		log:    log,
		server: server,
		router: router,
	}

	return app
}

func (a *Application) RegisterHandlers() {
	a.router.GET("/ping", a.Handlers.Status.Ping())

	authRoute := a.router.Group("/authhandler")
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
		a.log.Error(fmt.Sprintf("start application: %v", err))
	}
}

func (a *Application) gracefulShutdown() {
	ctx := context.Background()

	signalCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-signalCtx.Done()
	a.log.Info("shutting down server...")

	const waitingTime = 5 * time.Second

	shutdownCtx, cancel := context.WithTimeout(context.Background(), waitingTime)
	defer cancel()

	if err := a.server.Shutdown(shutdownCtx); err != nil {
		a.log.Error("failed to gracefully shutdown server", "error", err.Error())
	}
}
