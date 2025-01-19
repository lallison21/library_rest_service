package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lallison21/library_rest_service/internal/models"
)

type AuthRepo struct {
	pg *pgxpool.Pool
}

func NewAuth(pg *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		pg: pg,
	}
}

func (r *AuthRepo) Register(ctx context.Context, newUser *models.UserDAO) (int, error) {
	_ = ctx
	_ = newUser

	return 1, nil
}
