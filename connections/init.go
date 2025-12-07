package connections

import "database/sql"

const SQLITE = "sqlite"
const POSTGRES = "postgres"

type Connection interface {
	Connect() (*sql.DB, string, error)
}
