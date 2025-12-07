package base

import (
	"database/sql"
	"fmt"
)

func (manager Manager) Read(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := manager.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query '%s': %v", query, err)
	}
	return rows, nil
}

func (manager Manager) Write(query string, args ...interface{}) error {
	_, err := manager.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error writing query '%s': %v", query, err)
	}
	return nil
}

// ExecuteInTransaction will safely execute multiple queries
// execution in done in a func which will return error if anything gone wrong
func (manager Manager) ExecuteInTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := manager.DB.Begin()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// PrepareAndExecute is optimum for bulk insert
func (manager Manager) PrepareAndExecute(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := manager.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
