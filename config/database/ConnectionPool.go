package database

import (
	"database/sql"
)

type ConnectionPool struct {
	datasource DB
}

var connectionPool *ConnectionPool

func Instance() *ConnectionPool {
	if connectionPool == nil {
		connectionPool = &ConnectionPool{}
		connectionPool.Init("mysql")
	}
	return connectionPool
}

func (cp *ConnectionPool) Init(driver string) {
	if driver == "mysql" {
		cp.datasource = &MySQL{}
		cp.datasource.Init()
	}
}

func (cp *ConnectionPool) Connection() *sql.DB {
	return cp.datasource.Connection()
}

func (cp *ConnectionPool) Close() {
	cp.datasource.Close()
}
