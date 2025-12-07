package psql

import "fmt"

type Dialect struct{}

func (dialect Dialect) Placeholder(nos int) string {
	return fmt.Sprintf("$%d", nos)
}

func (dialect Dialect) Quote(identifier string) string {
	return `"` + identifier + `"`
}

func (dialect Dialect) LastInsertIDReturning() bool {
	return true
}
