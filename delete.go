package main

import (
	"fmt"
)

func (manager Manager) Delete(tableName string, where map[string]string) error {

	if !manager.CheckTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	// Build WHERE clause (start placeholders at 1)
	whereClause, whereValues := BuildWhereClause(manager.Dialect, where, 1)

	// Quote table name
	tbl := manager.Dialect.Quote(tableName)

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		tbl,
		whereClause,
	)

	return manager.Write(query, whereValues...)
}
