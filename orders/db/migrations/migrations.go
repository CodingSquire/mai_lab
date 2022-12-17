package migrations

import (
	"database/sql"
	"log"

	"github.com/lucsky/cuid"
)

type MigrationFunc func(db *sql.DB) error

// PerformMigration makes migration if not exists
// TODO: count hash of function innner
func PerformMigration(db *sql.DB, name string, migration MigrationFunc) (err error) {
	var id string
	_ = db.QueryRow("SELECT id FROM table_migrations WHERE name=$1", name).Scan(&id)
	if id == "" {
		_, err = db.Exec("INSERT INTO table_migrations (id, name) VALUES ($1, $2)", cuid.New(), name)
		if err != nil {
			return
		}
		err = migration(db)
	} else {
		log.Printf("%s already used: %s\n", name, id)
	}

	return
}
