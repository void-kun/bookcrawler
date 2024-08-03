package sql

import (
	"errors"
	"fmt"

	"bws/pkg/trace"

	"github.com/jmoiron/sqlx"
)

func GetSqlConnectionString(dataSourceType string, host string, port string, username string, password string, database string) (connStr string, err error) {
	if dataSourceType == "mysql" {
		connStr = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	} else if dataSourceType == "postgres" {
		connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, username, password, database, "disable")
	} else {
		err = errors.New(dataSourceType + "is not supported")
		return connStr, trace.TraceError(err)
	}
	return connStr, nil
}

func GetSqlConnection(dataSourceType string, host string, port string, username string, password string, database string) (db *sqlx.DB, err error) {
	connStr, err := GetSqlConnectionString(dataSourceType, host, port, username, password, database)
	if err != nil {
		return nil, trace.TraceError(err)
	}

	db, err = sqlx.Open(dataSourceType, connStr)
	if err != nil {
		return db, trace.TraceError(err)
	}

	return db, nil
}
