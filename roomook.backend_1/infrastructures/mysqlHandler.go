package infrastructures

import (
	"database/sql"
	"fmt"

	"jacob.de/roomook/roomook.backend/interfaces"
)

type MysqlHandler struct {
	Conn *sql.DB
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		return new(MySQLRow), err
	}
	return rows, nil
}

func (handler *MysqlHandler) PreparedQuery(statement string, args ...interface{}) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(MySQLRow), err
	}
	return rows, nil
}

func (handler *MysqlHandler) PreparedQueryResult(statement string, args ...interface{}) (interfaces.IResult, error) {
	stmt, err := handler.Conn.Prepare(statement)
	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
	}
	result, err := stmt.Exec(args...)

	if err != nil {
		return result, err
	}
	return result, nil
}

type MySQLRow struct {
	Rows *sql.Rows
}

func (r MySQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}

func (r MySQLRow) Close() error {
	return r.Rows.Close()
}
