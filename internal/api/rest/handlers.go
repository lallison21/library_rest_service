package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/api"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	log     *slog.Logger
	service api.Service
}

func NewHandler(log *slog.Logger, service api.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Info("handling ping request")
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, "pong")
	}
}
