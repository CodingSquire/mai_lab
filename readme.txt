curl -X POST -d '{"name":"user5678"}' localhost:8000/create
-->// {"id":"b40a9f54-bad4-4b98-9350-2a46507b43db","name":"user5678"}

curl -X POST localhost:8000/read?uid=b40a9f54-bad4-4b98-9350-2a46507b43db
-->// {"id":"b40a9f54-bad4-4b98-9350-2a46507b43db","name":"user5678"}

