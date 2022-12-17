package migrations

import (
	"database/sql"
	"github.com/lucsky/cuid"
)

type MigrationFunc func(db *sql.DB) error

// PerformMigration makes migration if not exists
// TODO: count hash of function innner
func PerformMigration(db *sql.DB, name string, migration MigrationFunc) (err error) {
	stmt, err := db.Prepare("select id, name from table_migrations where id = ?")
	if err != nil {
		return
	}

	err = stmt.QueryRow(name).Err()
	if err != nil {
		stmt, err = db.Prepare("insert into table_migrations (id, name) value (?, ?)")
		if err != nil {
			return
		}
		_, err = stmt.Exec(cuid.New(), name)
		if err != nil {
			return
		}
		err = migration(db)
	}

	return
}
