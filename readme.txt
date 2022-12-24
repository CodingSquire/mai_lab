docker-compose build
docker-compose up

migration:
docker run -v $/migrations:/migrations --network mai_lab_go_app migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@postgres_container:5432/postgres?sslmode=disable up 2

serv:
postgres all


create:
curl  -X POST -d '{"name":"user5671", "email": "oleg@kovinev.ru", "Phone":"9167743904"}' localhost:8000/create

Check:
curl -X POST localhost:8000/read?uuid=901a8f51-2c7d-4fb9-9749-2ddc05d0f7c6

Find:
curl -X POST localhost:8000/search?name=user5671

Delete:
curl -X POST localhost:8000/delete?uuid=901a8f51-2c7d-4fb9-9749-2ddc05d0f7c6


