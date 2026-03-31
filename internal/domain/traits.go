// setup:feature:graph

package domain

import "database/sql"

// ToNullString converts a string to sql.NullString. Empty strings are treated as null.
func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
