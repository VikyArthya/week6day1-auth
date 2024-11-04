package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConn() (*sql.DB, error) {
	connStr := "user=postgres password=superadmin dbname=go_auth sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
