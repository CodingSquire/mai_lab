createdb:
	docker exec -it pg_container createdb --username=root --owner=root users_service

dropdb:
	docker exec -it pg_container dropdb --username=root --owner=root users_service

migrateup:
	migrate -path ./migrations -database "postgresql://root:root@localhost:5432/users_service?sslmode=disable" -verbose up

migratedown:
	migrate -path ./migrations -database "postgresql://root:root@localhost:5432/users_service?sslmode=disable" -verbose down