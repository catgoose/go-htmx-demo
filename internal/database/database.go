// setup:feature:database

// Package database provides database connection management.
// Use OpenURL for app databases (parses scheme to pick driver + dialect).
// Use OpenSQLite for framework-internal stores (error traces, session settings).
package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"catgoose/dothog/internal/database/dialect"

	"github.com/jmoiron/sqlx"
)

// OpenURL opens a database connection from a URL string. The scheme determines
// the driver and dialect:
//
//	postgres://user:pass@host:5432/dbname?sslmode=disable
//	sqlite:///path/to/db.sqlite  or  sqlite:///:memory:
//	sqlserver://user:pass@host:1433?database=dbname
//
// Returns the raw *sql.DB and the matching Dialect for SQL generation.
func OpenURL(ctx context.Context, dsn string) (*sql.DB, dialect.Dialect, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("parse database URL: %w", err)
	}

	var engine dialect.Engine
	switch u.Scheme {
	case "postgres", "postgresql":
		engine = dialect.Postgres
	case "sqlite", "sqlite3":
		engine = dialect.SQLite
	case "sqlserver", "mssql":
		engine = dialect.MSSQL
	default:
		return nil, nil, fmt.Errorf("unsupported database scheme: %q", u.Scheme)
	}

	d, err := dialect.New(engine)
	if err != nil {
		return nil, nil, err
	}

	driverName := string(engine)
	connectStr := dsn
	if engine == dialect.SQLite {
		connectStr = u.Host + u.Path
		if connectStr == "" {
			connectStr = u.Opaque
		}
		driverName = "sqlite3"
	}

	db, err := sql.Open(driverName, connectStr)
	if err != nil {
		return nil, nil, fmt.Errorf("open %s: %w", engine, err)
	}
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, nil, fmt.Errorf("ping %s: %w", engine, err)
	}
	return db, d, nil
}

// OpenSQLite opens a SQLite database at the given path with standard settings.
// Used for framework-internal stores (error traces, session settings) that are
// always SQLite regardless of the app's primary database.
func OpenSQLite(ctx context.Context, dbPath string) (*sqlx.DB, error) {
	if dbPath != ":memory:" {
		if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
			return nil, fmt.Errorf("failed to create database directory: %w", err)
		}
	}

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to enable WAL mode: %w", err)
	}
	if _, err := db.Exec("PRAGMA busy_timeout=30000"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to set busy timeout: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping SQLite database: %w", err)
	}

	return db, nil
}

