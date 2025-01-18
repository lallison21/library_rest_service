package auth_handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lallison21/library_rest_service/internal/api"
	"github.com/lallison21/library_rest_service/internal/models"
	"log/slog"
	"net/http"
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
	return func(c *gin.Context) {
		var newUser models.CreateUser

		if err := c.BindJSON(&newUser); err != nil {
			var validation validator.ValidationErrors
			if errors.As(err, &validation) {
				for _, f := range validation {
					h.logging.Error("[Register] struct filed", "name:", f.Field(), "tag:", f.Tag())
				}
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := c.Request.Context()

		userId, err := h.service.Register(ctx, newUser.MapToDTO())
		if err != nil {
			h.logging.Error(fmt.Sprintf("[Register] [Register] create user: %v", err))
		}

		c.JSON(http.StatusOK, gin.H{"userId": userId})
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
