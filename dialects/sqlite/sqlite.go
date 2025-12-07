package sqlite

type Dialect struct{}

func (dialect Dialect) Placeholder(no int) string {
	return "?"
}

func (dialect Dialect) Quote(identifier string) string {
	return `"` + identifier + `"`
}

func (dialect Dialect) LastInsertIDReturning() bool {
	return false
}
