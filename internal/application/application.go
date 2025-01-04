package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/config/logging"
	"log/slog"
)

type Application struct {
	cfg     *config.Config
	logging *slog.Logger
}

func New() *Application {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	logger := logging.New(cfg.Logging)

	return &Application{
		cfg:     &cfg,
		logging: logger,
	}
}

func (a *Application) Run() {
	router := gin.New()

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}
	gin.SetMode(a.cfg.Server.GinMode)

	router.GET("/ping", func(c *gin.Context) {
		a.logging.Debug("pong")
		c.JSON(200, "pong")
	})

	addr := fmt.Sprintf("%s:%s", a.cfg.Server.Host, a.cfg.Server.Port)
	a.logging.Info(fmt.Sprintf("Listening on %s", addr))

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
