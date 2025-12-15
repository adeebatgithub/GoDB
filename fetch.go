package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/adeebatgithub/leaform/dialects"
)

func RowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))

		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := values[i]
			m[colName] = val
		}
		data = append(data, m)
	}
	return data, nil
}

func RowsToStringSlice(rows *sql.Rows) ([]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var data []interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))

		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		for i := range cols {
			val := values[i]
			data = append(data, val)
		}
	}
	return data, nil
}

func BuildWhereClause(d dialects.Dialect, conditions map[string]string, startIndex int) (string, []interface{}) {
	var whereClauses []string
	var values []interface{}

	i := startIndex
	for key, value := range conditions {
		col := d.Quote(key)
		ph := d.Placeholder(i)

		whereClauses = append(whereClauses, fmt.Sprintf("%s = %s", col, ph))
		values = append(values, value)

		i++
	}

	return strings.Join(whereClauses, " AND "), values
}

func (manager Manager) FetchAll(tableName, orderBy string, desc bool) ([]map[string]interface{}, error) {
	if !manager.CheckTableExists(tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}

	tbl := manager.Dialect.Quote(tableName)
	col := manager.Dialect.Quote("id")

	if orderBy != "" {
		col = manager.Dialect.Quote(orderBy)
	}

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s", tbl, col)
	if desc {
		query += " DESC"
	}

	rows, err := manager.Read(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return RowsToMap(rows)
}

func (manager Manager) FetchWhere(tableName string, where map[string]string, orderBy string, desc bool) ([]map[string]interface{}, error) {
	if !manager.CheckTableExists(tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}

	tbl := manager.Dialect.Quote(tableName)
	conditions, values := BuildWhereClause(manager.Dialect, where, 1)

	col := manager.Dialect.Quote("id")
	if orderBy != "" {
		col = manager.Dialect.Quote(orderBy)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s ORDER BY %s", tbl, conditions, col)
	if desc {
		query += " DESC"
	}

	rows, err := manager.Read(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return RowsToMap(rows)
}

func (manager Manager) FetchCol(tableName, colName, orderBy string, desc bool) ([]interface{}, error) {
	if !manager.CheckTableExists(tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}

	tbl := manager.Dialect.Quote(tableName)
	col := manager.Dialect.Quote(colName)

	ordCol := manager.Dialect.Quote("id")
	if orderBy != "" {
		ordCol = manager.Dialect.Quote(orderBy)
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s", col, tbl, ordCol)
	if desc {
		query += " DESC"
	}

	rows, err := manager.Read(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return RowsToStringSlice(rows)
}

func (manager Manager) FetchByID(tableName, id string) ([]map[string]interface{}, error) {
	tbl := manager.Dialect.Quote(tableName)
	col := manager.Dialect.Quote("id")
	ph := manager.Dialect.Placeholder(1)

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=%s", tbl, col, ph)

	rows, err := manager.Read(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return RowsToMap(rows)
}
