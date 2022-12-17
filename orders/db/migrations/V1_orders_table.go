package migrations

import (
	"database/sql"
	"log"
)

func V1(db *sql.DB) (err error) {
	createTable := `
	CREATE TABLE orders (
		id varchar(20) NOT NULL PRIMARY KEY,
		userId varchar(255) NOT NULL,
		item varchar(255) NOT NULL,
		adress text,
		count integer NOT NULL,
		createdAt timestamp NOT NULL DEFAULT NOW(),
		updatedAt timestamp NOT NULL DEFAULT NOW()
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Println("createTale")
		return
	}

	return
}
