package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Init initializate database and run migrations
func Init() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("COCKROACH_DB"))

	err = db.Ping()
	return db, err
}

func InitWithoutPing() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("COCKROACH_DB"))
	return db, err
}
