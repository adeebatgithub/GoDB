package psql

import (
	"fmt"
)

func (dialect Dialect) PrimaryKey() string {
	return "SERIAL PRIMARY KEY"
}

func (dialect Dialect) VarCharField(length int, unique bool, notNull bool) string {
	query := fmt.Sprintf("VARCHAR(%d)", length)
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
	query := "BIGINT"
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
	query := "BOOLEAN"
	if notNull {
		query += " NOT NULL"
	}
	query += fmt.Sprintf(" DEFAULT %t", defaultVal)
	return query
}

func (dialect Dialect) DateField(notNull bool) string {
	query := "DATE"
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) DateTimeField(onAdd bool) string {
	query := "TIMESTAMP"

	if onAdd {
		query += " DEFAULT CURRENT_TIMESTAMP"
	}

	return query
}

func (dialect Dialect) TimestampField(defaultNow, notNull bool) string {
	query := "TIMESTAMP"
	if notNull {
		query += " NOT NULL"
	}
	if defaultNow {
		query += " DEFAULT CURRENT_TIMESTAMP"
	}
	return query
}

func (dialect Dialect) DecimalField(precision, scale int, notNull bool) string {
	query := fmt.Sprintf("DECIMAL(%d, %d)", precision, scale)
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func (dialect Dialect) ForeignKeyField(refTable, refColumn string, notNull bool, onDelete string) string {
	query := fmt.Sprintf("INTEGER REFERENCES %s(%s)", refTable, refColumn)
	if onDelete != "" {
		query += " ON DELETE " + onDelete
		notNull = false
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}
