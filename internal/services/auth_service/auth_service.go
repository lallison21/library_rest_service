package auth_service

import (
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
