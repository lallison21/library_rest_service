package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/api"
)

type StatusHandler struct {
	log     *slog.Logger
	service api.StatusService
}

func NewStatus(log *slog.Logger, service api.StatusService) *StatusHandler {
	return &StatusHandler{
		log:     log,
		service: service,
	}
}

func (h *StatusHandler) Ping() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		h.log.Info("[Ping] start ping handler")

		pong, err := h.service.Ping()
		if err != nil {
			h.log.Error("[Ping] Iternal Server Error:" + err.Error())
			ginCtx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}

		h.log.Info("[Ping] request:" + pong)
		ginCtx.JSON(http.StatusOK, pong)
	}
}
