package database

import (
	"database/sql"
	"fmt"
	"strings"
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

func BuildWhereClause(conditions map[string]string, startindex int) (string, []interface{}) {
	var whereClauses []string
	var values []interface{}

	i := startindex
	for key, value := range conditions {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, value)
		i++
	}

	whereCondition := strings.Join(whereClauses, " AND ")

	return whereCondition, values
}

func FetchAll(db *sql.DB, tableName string, orderBy string, desc bool) ([]map[string]interface{}, error) {
	if !CheckTableExists(db, tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	if orderBy == "" {
		query += " ORDER BY id"
	} else {
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
	}
	if desc {
		query += " DESC"
	}
	rows, err := Read(db, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data, err := RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FetchOne(db *sql.DB, tableName string, where map[string]string, orderBy string, desc bool) ([]map[string]interface{}, error) {
	if !CheckTableExists(db, tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}
	conditions, values := BuildWhereClause(where, 1)
	query := fmt.Sprintf("SELECT * FROM %s where %s", tableName, conditions)
	if orderBy == "" {
		query += " ORDER BY id"
	} else {
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
	}
	if desc {
		query += " DESC"
	}
	rows, err := Read(db, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FetchCol(db *sql.DB, tableName, colName string, orderBy string, desc bool) ([]interface{}, error) {
	if !CheckTableExists(db, tableName) {
		return nil, fmt.Errorf("table `%s` does not exists", tableName)
	}
	query := fmt.Sprintf("SELECT %s FROM %s", colName, tableName)
	if orderBy == "" {
		query += " ORDER BY id"
	} else {
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
	}
	if desc {
		query += " DESC"
	}
	row, err := Read(db, query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	data, err := RowsToStringSlice(row)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FetchByID(db *sql.DB, tableName string, id int) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", tableName)
	rows, err := Read(db, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data, err := RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}
