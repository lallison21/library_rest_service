package auth_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/api"
	"log/slog"
)

type Handler struct {
	service api.AuthService
	logging *slog.Logger
}

func New(service api.AuthService, logging *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logging: logging,
	}
}

func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
