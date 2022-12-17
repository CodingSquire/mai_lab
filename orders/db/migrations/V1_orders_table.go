package migrations

import "database/sql"

func V1(db *sql.DB) (err error) {
	createTable := `
	create table orders (
		id varchar(20) not null primary key,
		userId varchar(255) not null,
		item varchar(255) not null,
		adress text,
		count integer not null,
		createdAt timestamp default DEFAULT NOW(),
		updatedAt timestamp default DEFAULT NOW(),
	)
	`
	_, err = db.Exec(createTable)
	if err != nil {
		return
	}
	
	createFunction := `
	CREATE OR REPLACE FUNCTION OrderUpdateAt
	RETURNS TRIGGER AS $$
	BEGIN
	  NEW.updatedAt = NOW();
	  RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	`
	_, err = db.Exec(createFunction)
	if err != nil {
		return
	}

	createTrigger := `
	CREATE TRIGGER set_timestamp
	BEFORE UPDATE ON orders
	FOR EACH ROW
	EXECUTE PROCEDURE trigger_set_timestamp();
	`
	_, err = db.Exec(createTrigger)
	if err != nil {
		return
	}

	return
}
