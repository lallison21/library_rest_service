package api

import (
	"context"
	"github.com/lallison21/library_rest_service/internal/models"
)

type StatusService interface {
	Ping() (string, error)
}

type AuthService interface {
	Register(ctx context.Context, user *models.UserDTO) (int, error)
}
