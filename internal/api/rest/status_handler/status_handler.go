package status_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/api"
	"log/slog"
	"net/http"
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
	return func(c *gin.Context) {
		h.log.Info("[Ping] start ping handler")

		pong, err := h.service.Ping()
		if err != nil {
			h.log.Error("[Ping] Iternal Server Error:" + err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}

		h.log.Info("[Ping] request:" + pong)
		c.JSON(http.StatusOK, pong)
	}
}
