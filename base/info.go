package base

func (manager Manager) GetTableNames() ([]string, error) {
	rows, err := manager.Read(manager.Dialect.TableNames())
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

func (manager Manager) CheckTableExists(tableName string) bool {
	tables, err := manager.GetTableNames()
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

func (manager Manager) GetColumnNames(tableName string) ([]string, error) {
	rows, err := manager.Read(manager.Dialect.ColumnNames(tableName))
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

func (manager Manager) CheckColumnExists(tableName string, columnName string) bool {
	columns, err := manager.GetColumnNames(tableName)
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
