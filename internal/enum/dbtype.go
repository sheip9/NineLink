package enum

type DBType string

const (
	MySQL    = DBType("mysql")
	Postgres = DBType("postgres")
)
