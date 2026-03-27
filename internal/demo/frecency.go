// setup:feature:demo

package demo

import (
	"context"
)

// PageVisit tracks visit frequency and recency for a page.
type PageVisit struct {
	Path      string
	Title     string
	Visits    int
	LastVisit string
	Score     float64
}

// initFrecency creates the page_visits table.
func (d *DB) initFrecency() error {
	_, err := d.db.Exec(`CREATE TABLE IF NOT EXISTS page_visits (
		session_id TEXT    NOT NULL,
		path       TEXT    NOT NULL,
		title      TEXT    NOT NULL DEFAULT '',
		visits     INTEGER NOT NULL DEFAULT 1,
		last_visit TEXT    NOT NULL DEFAULT (datetime('now')),
		PRIMARY KEY (session_id, path)
	)`)
	return err
}

// RecordVisit upserts a page visit for a session.
func (d *DB) RecordVisit(ctx context.Context, sessionID, path, title string) error {
	_, err := d.db.ExecContext(ctx, `
		INSERT INTO page_visits (session_id, path, title, visits, last_visit)
		VALUES (?, ?, ?, 1, datetime('now'))
		ON CONFLICT (session_id, path) DO UPDATE SET
			visits = visits + 1,
			last_visit = datetime('now'),
			title = excluded.title
	`, sessionID, path, title)
	return err
}

// TopFrecent returns the top N most frecent pages for a session.
// Score is visits / (days_since_last_visit + 1), so recent frequent pages rank highest.
func (d *DB) TopFrecent(ctx context.Context, sessionID string, limit int) ([]PageVisit, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT path, title, visits, last_visit,
		       CAST(visits AS REAL) / (julianday('now') - julianday(last_visit) + 1) AS score
		FROM page_visits
		WHERE session_id = ?
		ORDER BY score DESC
		LIMIT ?
	`, sessionID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PageVisit
	for rows.Next() {
		var pv PageVisit
		if err := rows.Scan(&pv.Path, &pv.Title, &pv.Visits, &pv.LastVisit, &pv.Score); err != nil {
			return nil, err
		}
		results = append(results, pv)
	}
	return results, rows.Err()
}

// PopularPages returns the top N most popular pages across all sessions.
func (d *DB) PopularPages(ctx context.Context, limit int) ([]PageVisit, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT path, title,
		       SUM(visits) as total_visits,
		       MAX(last_visit) as latest,
		       CAST(SUM(visits) AS REAL) / (julianday('now') - julianday(MAX(last_visit)) + 1) AS score
		FROM page_visits
		GROUP BY path
		ORDER BY score DESC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PageVisit
	for rows.Next() {
		var pv PageVisit
		if err := rows.Scan(&pv.Path, &pv.Title, &pv.Visits, &pv.LastVisit, &pv.Score); err != nil {
			return nil, err
		}
		results = append(results, pv)
	}
	return results, rows.Err()
}
