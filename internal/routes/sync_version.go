// setup:feature:sync
package routes

import (
	"context"
)

// knownTables is the set of tables that support version-based sync.
// This is the single source of truth for SQL table name validation —
// no regex needed when we can enumerate the valid values.
var knownTables = map[string]bool{
	"Tasks":  true,
	"Items":  true,
	"People": true,
}

// VersionChecker looks up the current version of a row by table and ID.
type VersionChecker interface {
	CurrentVersion(ctx context.Context, table string, id int) (int, error)
}
