// setup:feature:postgres

package dialect

import "fmt"

// PostgresDialect implements Dialect for PostgreSQL.
type PostgresDialect struct{}

func (PostgresDialect) Engine() Engine    { return Postgres }
func (PostgresDialect) Pagination() string { return "LIMIT @Limit OFFSET @Offset" }
func (PostgresDialect) AutoIncrement() string {
	return "SERIAL PRIMARY KEY"
}
func (PostgresDialect) Now() string           { return "NOW()" }
func (PostgresDialect) TimestampType() string { return "TIMESTAMPTZ" }
func (PostgresDialect) StringType(maxLen int) string {
	return fmt.Sprintf("VARCHAR(%d)", maxLen)
}
func (PostgresDialect) VarcharType(maxLen int) string {
	return fmt.Sprintf("VARCHAR(%d)", maxLen)
}
func (PostgresDialect) IntType() string  { return "INTEGER" }
func (PostgresDialect) TextType() string { return "TEXT" }

func (PostgresDialect) CreateTableIfNotExists(table, body string) string {
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, body)
}

func (PostgresDialect) DropTableIfExists(table string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
}

func (PostgresDialect) CreateIndexIfNotExists(indexName, table, columns string) string {
	return fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", indexName, table, columns)
}

func (PostgresDialect) LastInsertIDQuery() string { return "" }
func (PostgresDialect) SupportsLastInsertID() bool { return false }

func (PostgresDialect) TableExistsQuery() string {
	return "SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1"
}

func (PostgresDialect) TableColumnsQuery() string {
	return "SELECT column_name AS name FROM information_schema.columns WHERE table_schema = 'public' AND table_name = $1 ORDER BY ordinal_position"
}
