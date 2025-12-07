package edit

import (
	"fmt"
)

func (manager Manager) DeleteByColumn(tableName, column string, value any) error {
	if !manager.CheckTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	tbl := manager.Dialect.Quote(tableName)
	col := manager.Dialect.Quote(column)
	ph := manager.Dialect.Placeholder(1)

	query := fmt.Sprintf("DELETE FROM %s WHERE %s=%s", tbl, col, ph)
	return manager.Write(query, value)
}
