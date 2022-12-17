package db

import (
	"database/sql"
	"orders/db/migrations"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// enshureMigrationTable enshures that table_migrations exists
func enshureMigrationTable(db *sql.DB) (err error) {
	sqlStmt :=`
	CREATE TABLE IF NOT EXISTS table_migrations (id varchar(255) not null primary key, name varchar(255) not null);
	`
	_, err = db.Exec(sqlStmt)

	return
}

// Init initializate database and run migrations
func Init() (db *sql.DB, err error) {
	db, err = sql.Open("pgx", os.Getenv("COCKROACH_DB"))
	if err != nil {
		return
	}

	err = enshureMigrationTable(db)
	if err != nil {
		return
	}

	err = migrations.PerformMigration(db, "V1", migrations.V1)

	return;
}
