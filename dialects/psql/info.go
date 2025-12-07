package psql

import "fmt"

func (dialect Dialect) TableNames() string {
	return "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';"
}

func (dialect Dialect) ColumnNames(tableName string) string {
	query := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name = '%s'", tableName)
	return query
}
