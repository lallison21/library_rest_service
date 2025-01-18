package services

import (
	"context"
	"github.com/lallison21/library_rest_service/internal/models"
)

type StatusRepo interface {
	Ping() (string, error)
}

type AuthRepo interface {
	Register(ctx context.Context, newUser *models.UserDAO) (int, error)
}

type PasswordUtils interface {
	GeneratePassword(password string) (string, error)
	ComparePassword(password, hash string) (bool, error)
}
