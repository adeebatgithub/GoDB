package dialects

type Dialect interface {
	Placeholder(nos int) string
	Quote(identifier string) string
	LastInsertIDReturning() bool

	TableNames() string
	ColumnNames(tableName string) string

	PrimaryKey() string
	ForeignKeyField(refTable, refColumn string, notNull bool, onDelete string) string
	VarCharField(length int, unique bool, notNull bool) string
	TextField(notNull bool) string
	IntegerField(unique, notNull bool) string
	BigIntField(unique, notNull bool) string
	DecimalField(precision, scale int, notNull bool) string
	BooleanField(defaultVal bool, notNull bool) string
	DateField(notNull bool) string
	DateTimeField(onAdd bool) string
	TimestampField(defaultNow, notNull bool) string
}
