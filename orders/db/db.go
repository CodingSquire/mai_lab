package db

import (
	"database/sql"
	"orders/db/migrations"
	"os"
)

// enshureMigrationTable enshures that table_migrations exists
func enshureMigrationTable(db *sql.DB) (err error) {
	sqlStmt :=`
	create if not exists table_migrations (id text not null primary key, name text not null);
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
