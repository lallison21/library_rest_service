package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"github.com/lallison21/library_rest_service/internal/services"
	"log/slog"
	"net/http"
)

func New(cfg *config.Server, logging *slog.Logger, service services.Service) *http.Server {
	router := gin.New()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}
	gin.SetMode(cfg.GinMode)

	handler := NewHandler(logging, service)
	router.GET("/ping", handler.Ping())

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
