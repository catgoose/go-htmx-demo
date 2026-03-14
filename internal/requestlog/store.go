// Package requestlog provides per-request log capture with promote-on-error
// semantics. Each request buffers its slog records locally; only when an error
// occurs is the buffer promoted to a SQLite-backed Store for later retrieval.
package requestlog

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	dbrepo "catgoose/dothog/internal/database/repository"
	"catgoose/dothog/internal/database/schema"

	"github.com/jmoiron/sqlx"
)

// Entry is a single captured log record.
type Entry struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"msg"`
	Attrs   string    `json:"attrs,omitempty"`
}

// Buffer is a per-request log buffer stored in the request context.
// It is not thread-safe — each request is handled by a single goroutine.
type Buffer struct {
	Entries []Entry
}

type bufferKey struct{}

// NewBufferContext returns a new context with an empty Buffer attached.
func NewBufferContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, bufferKey{}, &Buffer{})
}

// GetBuffer retrieves the per-request Buffer from the context, or nil.
func GetBuffer(ctx context.Context) *Buffer {
	buf, _ := ctx.Value(bufferKey{}).(*Buffer)
	return buf
}

var tableName = schema.ErrorTracesTable.Name

// Store is a SQLite-backed store of error request log entries.
// Only requests that encounter errors are promoted here.
type Store struct {
	db *sqlx.DB
}

// NewStore creates a Store backed by the given database connection.
func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

// Promote persists a per-request buffer to the database. This should only
// be called when the request resulted in an error.
func (s *Store) Promote(requestID string, entries []Entry) {
	if len(entries) == 0 {
		return
	}
	data, err := json.Marshal(entries)
	if err != nil {
		return
	}
	insertCols := schema.ErrorTracesTable.InsertColumns()
	query := dbrepo.InsertInto(tableName, insertCols...)
	now := dbrepo.GetNow()
	_, _ = s.db.Exec(query,
		sql.Named("RequestID", requestID),
		sql.Named("Entries", string(data)),
		sql.Named("CreatedAt", now),
		sql.Named("UpdatedAt", now),
	)
}

// errorTraceRow maps to a row in the ErrorTraces table.
type errorTraceRow struct {
	Entries string `db:"Entries"`
}

// Get returns all captured entries for a request ID, or nil if not found.
func (s *Store) Get(requestID string) []Entry {
	w := dbrepo.NewWhere().And("RequestID = @RequestID", sql.Named("RequestID", requestID))
	query, args := dbrepo.NewSelect(tableName, "Entries").Where(w).Build()

	var row errorTraceRow
	err := s.db.Get(&row, query, args...)
	if err != nil {
		return nil
	}
	var entries []Entry
	if err := json.Unmarshal([]byte(row.Entries), &entries); err != nil {
		return nil
	}
	return entries
}

// StartCleanup runs a background goroutine that deletes entries older than ttl.
// It stops when ctx is cancelled.
func (s *Store) StartCleanup(ctx context.Context, ttl time.Duration, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.deleteOlderThan(ttl)
			}
		}
	}()
}

func (s *Store) deleteOlderThan(ttl time.Duration) {
	cutoff := time.Now().Add(-ttl)
	query := fmt.Sprintf("DELETE FROM %s WHERE CreatedAt < @Cutoff", tableName)
	_, _ = s.db.Exec(query, sql.Named("Cutoff", cutoff))
}
