package database

import "database/sql"

type DB interface {
	Init()
	Connection() *sql.DB
	Close()
}
