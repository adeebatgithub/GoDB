package database

import (
	"database/sql"
	"fmt"
)

func GetTableNames(db *sql.DB) ([]string, error) {
	rows, err := Read(db, "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tableNames, nil
}

func CheckTableExists(db *sql.DB, tableName string) bool {
	tables, err := GetTableNames(db)
	if err != nil {
		return false
	}
	for _, table := range tables {
		if tableName == table {
			return true
		}
	}
	return false
}

func GetColumnNames(db *sql.DB, tableName string) ([]string, error) {
	query := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name = '%s'", tableName)
	rows, err := Read(db, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columnNames []string
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return nil, err
		}
		columnNames = append(columnNames, columnName)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columnNames, nil
}

func CheckColumnExists(db *sql.DB, tableName string, columnName string) bool {
	columns, err := GetColumnNames(db, tableName)
	if err != nil {
		return false
	}
	for _, column := range columns {
		if columnName == column {
			return true
		}
	}
	return false
}
