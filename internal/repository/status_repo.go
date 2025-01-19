package repository

import "github.com/jackc/pgx/v5/pgxpool"

type StatusRepo struct {
	pg *pgxpool.Pool
}

func NewStatus(pg *pgxpool.Pool) *StatusRepo {
	return &StatusRepo{
		pg: pg,
	}
}

func (r *StatusRepo) Ping() (string, error) {
	return "pong", nil
}
