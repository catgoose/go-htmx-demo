// setup:feature:session_settings

package repository

import (
	"context"
	"database/sql"
	"fmt"

	dbrepo "catgoose/harmony/internals/database/repository"
	"catgoose/harmony/internals/domain"
)

// SessionSettingsRepository defines operations for session settings data access.
type SessionSettingsRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*domain.SessionSettings, error)
	Upsert(ctx context.Context, s *domain.SessionSettings) error
	Touch(ctx context.Context, uuid string) error
	DeleteStale(ctx context.Context, days int) (int64, error)
}

// sessionSettingsRepository implements SessionSettingsRepository.
type sessionSettingsRepository struct {
	repo *dbrepo.RepoManager
}

// NewSessionSettingsRepository creates a new SessionSettingsRepository.
func NewSessionSettingsRepository(repo *dbrepo.RepoManager) SessionSettingsRepository {
	return &sessionSettingsRepository{repo: repo}
}

// GetByUUID returns settings for the given session UUID, or nil if not found.
func (r *sessionSettingsRepository) GetByUUID(ctx context.Context, uuid string) (*domain.SessionSettings, error) {
	var s domain.SessionSettings
	err := r.repo.GetDB().GetContext(ctx,
		&s,
		"SELECT Id, SessionUUID, Theme, UpdatedAt FROM SessionSettings WHERE SessionUUID = ?",
		uuid,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get session settings: %w", err)
	}
	return &s, nil
}

// Upsert inserts or updates session settings by SessionUUID.
func (r *sessionSettingsRepository) Upsert(ctx context.Context, s *domain.SessionSettings) error {
	existing, err := r.GetByUUID(ctx, s.SessionUUID)
	if err != nil {
		return err
	}
	if existing != nil {
		_, err = r.repo.GetDB().ExecContext(ctx,
			`UPDATE SessionSettings SET Theme = ?, UpdatedAt = CURRENT_TIMESTAMP
			 WHERE SessionUUID = ?`,
			s.Theme,
			s.SessionUUID,
		)
		if err != nil {
			return fmt.Errorf("update session settings: %w", err)
		}
		return nil
	}
	_, err = r.repo.GetDB().ExecContext(ctx,
		`INSERT INTO SessionSettings (SessionUUID, Theme, CreatedAt, UpdatedAt)
		 VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		s.SessionUUID,
		s.Theme,
	)
	if err != nil {
		return fmt.Errorf("insert session settings: %w", err)
	}
	return nil
}

// Touch updates UpdatedAt for the given session UUID.
func (r *sessionSettingsRepository) Touch(ctx context.Context, uuid string) error {
	_, err := r.repo.GetDB().ExecContext(ctx,
		"UPDATE SessionSettings SET UpdatedAt = CURRENT_TIMESTAMP WHERE SessionUUID = ?",
		uuid,
	)
	if err != nil {
		return fmt.Errorf("touch session settings: %w", err)
	}
	return nil
}

// DeleteStale removes session settings rows not updated in the given number of days.
func (r *sessionSettingsRepository) DeleteStale(ctx context.Context, days int) (int64, error) {
	res, err := r.repo.GetDB().ExecContext(ctx,
		"DELETE FROM SessionSettings WHERE UpdatedAt < datetime('now', ?)",
		fmt.Sprintf("-%d days", days),
	)
	if err != nil {
		return 0, fmt.Errorf("delete stale session settings: %w", err)
	}
	return res.RowsAffected()
}
