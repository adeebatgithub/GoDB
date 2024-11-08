package database

import (
	"database/sql"
	"fmt"
	"strings"
)

func CreateTable(db *sql.DB, tableName string, columns map[string]string) error {
	var columnDefinitions []string
	for columnName, dataType := range columns {
		columnDefinitions = append(columnDefinitions, fmt.Sprintf("%s %s", columnName, dataType))
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, strings.Join(columnDefinitions, ","))
	err := Write(db, query)
	if err != nil {
		return err
	}
	return nil
}

func Insert(db *sql.DB, tableName string, columns map[string]string) error {
	if !CheckTableExists(db, tableName) {
		return fmt.Errorf("table %s does not exist", tableName)
	}
	var columnNames []string
	var placeholders []string
	var values []interface{}

	i := 1
	for col, val := range columns {
		columnNames = append(columnNames, col)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, val)
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(columnNames, ", "),
		strings.Join(placeholders, ", "),
	)

	err := Write(db, query, values...)
	return err
}

func Update(db *sql.DB, tableName string, columns, where map[string]string) error {
	var setClauses []string
	var values []interface{}

	i := 1
	for col, val := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i))
		values = append(values, val)
		i++
	}

	setClause := strings.Join(setClauses, ", ")
	whereClause, whereValues := BuildWhereClause(where, i)
	values = append(values, whereValues...)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, setClause, whereClause)

	err := Write(db, query, values...)
	return err
}
