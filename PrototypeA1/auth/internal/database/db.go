package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open(
		"postgres", 
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	)
}