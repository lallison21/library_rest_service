package auth_service

import (
	"context"
	"fmt"
	"github.com/lallison21/library_rest_service/internal/models"
	"github.com/lallison21/library_rest_service/internal/services"
)

type AuthService struct {
	repo      services.AuthRepo
	passUtils services.PasswordUtils
}

func New(repo services.AuthRepo, passUtils services.PasswordUtils) *AuthService {
	return &AuthService{
		repo:      repo,
		passUtils: passUtils,
	}
}

func (s *AuthService) Register(ctx context.Context, newUser *models.UserDTO) (int, error) {
	hashedPassword, err := s.passUtils.GeneratePassword(newUser.Password)
	if err != nil {
		return -1, fmt.Errorf("[Register] generate password: %w", err)
	}
	newUser.Password = hashedPassword

	newUserId, err := s.repo.Register(ctx, newUser.MapToDAO())
	if err != nil {
		return -1, fmt.Errorf("[Register] register: %w", err)
	}

	return newUserId, nil
}
