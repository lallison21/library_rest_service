package auth_service

import (
	"context"
	"github.com/lallison21/library_rest_service/internal/models"
	"github.com/lallison21/library_rest_service/internal/services"
)

type AuthService struct {
	repo services.AuthRepo
}

func New(repo services.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, newUser *models.UserDAO) (int, error) {
	return s.repo.Register(ctx, newUser)
}
