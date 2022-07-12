package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn *sql.DB
}

func (m *MySQL) Init() {
	var err error
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	m.conn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	m.conn.SetMaxIdleConns(50)
	m.conn.SetMaxOpenConns(50)
	m.conn.SetConnMaxLifetime(time.Second)
}

func (m *MySQL) Connection() *sql.DB {
	return m.conn
}

func (m *MySQL) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}
