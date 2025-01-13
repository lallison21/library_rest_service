package status_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/api"
	"log/slog"
)

type Handler struct {
	log     *slog.Logger
	service api.StatusService
}

func New(log *slog.Logger, service api.StatusService) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
