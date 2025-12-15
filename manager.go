package main

import (
	"database/sql"

	"github.com/adeebatgithub/biscut/connections"
	"github.com/adeebatgithub/biscut/dialects"
	"github.com/adeebatgithub/biscut/dialects/psql"
	"github.com/adeebatgithub/biscut/dialects/sqlite"
)

type Manager struct {
	DB      *sql.DB
	Dialect dialects.Dialect
}

func DialectFromDriver(driver string) dialects.Dialect {
	switch driver {
	case connections.SQLITE:
		return sqlite.Dialect{}
	case connections.POSTGRES:
		return psql.Dialect{}
	default:
		return sqlite.Dialect{}
	}
}

func NewManager(connection connections.Connection) (*Manager, error) {
	db, driver, err := connection.Connect()
	if err != nil {
		return nil, err
	}
	return &Manager{DB: db, Dialect: DialectFromDriver(driver)}, nil
}
