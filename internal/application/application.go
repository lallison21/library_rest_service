package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/library_rest_service/internal/config"
)

type Application struct {
	cfg *config.Config
}

func New() *Application {
	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &Application{
		cfg: &cfg,
	}
}

func (a *Application) Run() {
	router := gin.New()

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}
	gin.SetMode(gin.DebugMode)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	addr := fmt.Sprintf("%s:%s", a.cfg.Server.Host, a.cfg.Server.Port)
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
