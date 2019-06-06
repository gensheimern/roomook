package interfaces

//go:generate mockgen -destination=./mocks/IDbHandler.go -package=mocks jacob.de/roombook/interfaces IDbHandler

type IDbHandler interface {
	Execute(statement string)
	Query(statement string) (IRow, error)
	PreparedQuery(statement string, args ...interface{}) (IRow, error)
	PreparedQueryResult(statement string, args ...interface{}) (IResult, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type IResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
