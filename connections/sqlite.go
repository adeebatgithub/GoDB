package connections

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Database string
}

func (config Sqlite) Connect() (*sql.DB, string, error) {
	db, err := sql.Open("sqlite", config.Database)
	if err != nil {
		return nil, "", err
	}
	return db, SQLITE, nil
}
