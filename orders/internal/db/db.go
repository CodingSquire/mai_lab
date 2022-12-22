package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Init initializate database and run migrations
func Init() (db *sql.DB, err error) {
	db, err = sql.Open("pgx", os.Getenv("COCKROACH_DB"))
	if err != nil {
		return
	}

	return
}
