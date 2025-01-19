package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lallison21/library_rest_service/internal/api"
	"github.com/lallison21/library_rest_service/internal/models"
)

type AuthHandler struct {
	service api.AuthService
	logging *slog.Logger
}

func NewAuth(service api.AuthService, logging *slog.Logger) *AuthHandler {
	return &AuthHandler{
		service: service,
		logging: logging,
	}
}

func (h *AuthHandler) Register() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		var newUser models.CreateUser

		if err := ginCtx.BindJSON(&newUser); err != nil {
			var validation validator.ValidationErrors
			if errors.As(err, &validation) {
				for _, f := range validation {
					h.logging.Error("[Register] struct filed", "name:", f.Field(), "tag:", f.Tag())
				}
			}

			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		ctx := ginCtx.Request.Context()

		userID, err := h.service.Register(ctx, newUser.MapToDTO())
		if err != nil {
			h.logging.Error(fmt.Sprintf("[Register] [Register] create user: %v", err))
		}

		ginCtx.JSON(http.StatusOK, gin.H{"userId": userID})
	}
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c
	}
}
