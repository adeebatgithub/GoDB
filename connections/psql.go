package connections

import (
	"database/sql"
	"fmt"
)

type PSQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SslMode  string
}

func (config PSQLConfig) Connect() (*sql.DB, string, error) {
	conf := fmt.Sprintf(
		"dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		config.Database, config.Host, config.Port, config.Username, config.Password, config.SslMode,
	)

	db, err := sql.Open("postgres", conf)
	if err != nil {
		return nil, "", err
	}

	err = db.Ping()
	if err != nil {
		return nil, "", err
	}

	return db, POSTGRES, nil
}
