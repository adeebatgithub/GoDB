package database

import (
	"fmt"
	"strconv"
)

func PrimaryKey() string {
	return "SERIAL PRIMARY KEY"
}

func VarCharField(length int, unique bool, notNull bool) string {
	query := "VARCHAR(" + strconv.Itoa(length) + ")"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func IntegerField(unique, notNull bool) string {
	query := "INTEGER"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func BigIntField(unique, notNull bool) string {
	query := "BIGINT"
	if unique {
		query += " UNIQUE"
	}
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func TextField(notNull bool) string {
	query := "TEXT"
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func BooleanField(defaultVal *bool, notNull bool) string {
	query := "BOOLEAN"
	if notNull {
		query += " NOT NULL"
	}
	if defaultVal != nil {
		query += fmt.Sprintf(" DEFAULT %t", *defaultVal)
	}
	return query
}

func DateField(notNull bool) string {
	query := "DATE"
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func TimestampField(defaultNow, notNull bool) string {
	query := "TIMESTAMP"
	if notNull {
		query += " NOT NULL"
	}
	if defaultNow {
		query += " DEFAULT CURRENT_TIMESTAMP"
	}
	return query
}

func DecimalField(precision, scale int, notNull bool) string {
	query := fmt.Sprintf("DECIMAL(%d, %d)", precision, scale)
	if notNull {
		query += " NOT NULL"
	}
	return query
}

func ForeignKeyField(refTable, refColumn string, notNull bool) string {
	query := fmt.Sprintf("INTEGER REFERENCES %s(%s)", refTable, refColumn)
	if notNull {
		query += " NOT NULL"
	}
	return query
}
