package auth_repo

import (
	"context"
	"github.com/lallison21/library_rest_service/internal/models"
)

type AuthRepo struct {
}

func New() *AuthRepo {
	return &AuthRepo{}
}

func (r *AuthRepo) Register(ctx context.Context, newUser *models.UserDAO) (int, error) {
	return 1, nil
}
