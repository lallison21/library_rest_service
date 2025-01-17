package status_repo

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pg *pgxpool.Pool
}

func New(pg *pgxpool.Pool) *Repository {
	return &Repository{
		pg: pg,
	}
}

func (r *Repository) Ping() (string, error) {
	return "pong", nil
}
