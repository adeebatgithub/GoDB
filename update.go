package main

import (
	"fmt"
	"strings"
)

func (manager Manager) Update(tableName string, columns, where map[string]string) error {
	if !manager.CheckTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	var setClauses []string
	var values []interface{}

	i := 1

	// Build SET col = placeholder
	for col, val := range columns {
		qcol := manager.Dialect.Quote(col)
		ph := manager.Dialect.Placeholder(i)

		setClauses = append(setClauses, fmt.Sprintf("%s = %s", qcol, ph))
		values = append(values, val)
		i++
	}

	setClause := strings.Join(setClauses, ", ")

	// WHERE clause using next placeholder index
	whereClause, whereValues := BuildWhereClause(manager.Dialect, where, i)
	values = append(values, whereValues...)

	// Quote table name
	tbl := manager.Dialect.Quote(tableName)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tbl, setClause, whereClause)

	return manager.Write(query, values...)
}
