package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lallison21/library_rest_service/internal/config/config"
)

func New(cfg config.Postgres, log *slog.Logger) *pgxpool.Pool {
	const waitingTime = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), waitingTime)
	defer cancel()

	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s "+
			"pool_max_conns=%d pool_max_conn_lifetime=%v pool_max_conn_idle_time=%v",
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
