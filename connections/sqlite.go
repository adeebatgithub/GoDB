package connections

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SqliteConfig struct {
	Database string
}

func (config SqliteConfig) Connect() (*sql.DB, string, error) {
	db, err := sql.Open("sqlite", config.Database)
	if err != nil {
		return nil, "", err
	}
	return db, SQLITE, nil
}
