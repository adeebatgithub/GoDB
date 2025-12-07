package sqlite

import (
	"fmt"
)

func (dialect Dialect) PrimaryKey() string {
	return "INTEGER PRIMARY KEY AUTOINCREMENT"
}

func (dialect Dialect) VarCharField(length int, unique bool, notNull bool) string {
	query := "TEXT"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) IntegerField(unique, notNull bool) string {
	query := "INTEGER"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) BigIntField(unique, notNull bool) string {
	query := "INTEGER"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) TextField(notNull bool) string {
	query := "TEXT"
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) BooleanField(defaultVal bool, notNull bool) string {
	query := "INTEGER"
	if notNull {
		query += " NOT NULL"
	}
	intVal := 0
	if defaultVal {
		intVal = 1
	}
	query += fmt.Sprintf(" DEFAULT %d", intVal)
	return query
}

func (dialect Dialect) DateField(notNull bool) string {
	query := "TEXT"
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) DateTimeField(onAdd bool) string {
	query := "TEXT"

	if onAdd {
		query += " DEFAULT (datetime('now'))"
	}

	return query
}

func (dialect Dialect) TimestampField(defaultNow, notNull bool) string {
	query := "TEXT"
	if notNull {
		query += " NOT NULL"
	}
	if defaultNow {
		query += " DEFAULT CURRENT_TIMESTAMP"
	}
	return query
}

func (dialect Dialect) DecimalField(precision, scale int, notNull bool) string {
	query := fmt.Sprintf("NUMERIC(%d, %d)", precision, scale)
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) ForeignKeyField(refTable, refColumn string, notNull bool) string {
	query := fmt.Sprintf("INTEGER REFERENCES %s(%s)", refTable, refColumn)
	if notNull {
		query += " NOT NULL"
	}
	return query
}
