package sqlite

import "fmt"

func (dialect Dialect) TableNames() string {
	return "SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'"
}

func (dialect Dialect) ColumnNames(tableName string) string {
	query := fmt.Sprintf("PRAGMA table_info(%s);", tableName)
	return query
}
