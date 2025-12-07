package base

import (
	"fmt"
	"strings"
)

type Table struct {
	TableName string
	Fields    map[string]string
}

func (manager Manager) CreateTable(table *Table) error {
	var columnDefinitions []string
	var tableName = table.TableName
	var columns = table.Fields
	for columnName, dataType := range columns {
		columnDefinitions = append(columnDefinitions, fmt.Sprintf("%s %s", columnName, dataType))
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, strings.Join(columnDefinitions, ","))
	err := manager.Write(query)
	if err != nil {
		return err
	}
	return nil
}
