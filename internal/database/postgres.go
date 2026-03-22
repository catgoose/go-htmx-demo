// setup:feature:postgres

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/catgoose/dio"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // register postgres driver
)

func openPostgresDB(ctx context.Context) (*sqlx.DB, error) {
	dsn, err := dio.Env("DATABASE_URL")
	if err != nil || dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is required for postgres")
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open Postgres database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping Postgres database: %w", err)
	}

	return db, nil
}
