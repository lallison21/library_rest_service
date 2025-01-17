package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"log/slog"
	"time"
)

func New(cfg config.Postgres, log *slog.Logger) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s pool_max_conns=%s pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password,
		cfg.SslMode,
		cfg.MaxConns,
		cfg.ConnMaxLifetime,
		cfg.ConnMaxIdleTime,
	)

	poolConfig, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		log.Error("parse postgres config", "err", err)
		panic(err)
	}

	connPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error("create postgres connection pool", "err", err)
		panic(err)
	}

	log.Info("connect to postgres successfully")
	return connPool
}
