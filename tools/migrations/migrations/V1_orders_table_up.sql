CREATE TABLE orders (
	id varchar(100) NOT NULL PRIMARY KEY,
	userId varchar(255) NOT NULL,
	item varchar(255) NOT NULL,
	adress text,
	count integer NOT NULL,
	createdAt timestamp NOT NULL DEFAULT NOW(),
	updatedAt timestamp NOT NULL DEFAULT NOW()
);
