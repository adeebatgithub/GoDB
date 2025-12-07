package main

import (
	"fmt"
	"strings"
)

func (manager Manager) Insert(tableName string, columns map[string]string) error {
	if !manager.CheckTableExists(tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}

	var colNames []string
	var placeholders []string
	var values []interface{}

	i := 1
	for col, val := range columns {
		colNames = append(colNames, manager.Dialect.Quote(col))
		placeholders = append(placeholders, manager.Dialect.Placeholder(i))
		values = append(values, val)
		i++
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		manager.Dialect.Quote(tableName),
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "),
	)

	return manager.Write(query, values...)
}
