# mai_lab

## Запуск
```
docker-compose build && docker-compose up
```

## Миграция БД
```
docker run -v $(pwd)/migrations:/migrations --network go_services_lab_go_app migrate/migrate -path=/migrations/ -database postgres://postgres:qweasd@postgres_container:5432/postgres?sslmode=disable up 2
```

## PgAdmin
По пути `http://localhost:5050`

## gRPC
Для генерации
```
protoc --proto_path=proto --go_out=plugins=grpc:pkg/user/proto --go_opt=paths=source_relative users.proto
```

## Задание №1 
### Сервис пользователей

- Написать сервис Users
- Сервис позволяет создавать/получать аккаунт пользователя
- Сервис должен хранить все данные в кеше

### Сервис заказов

- Написать сервис Orders
- Сервис позволяет создавать/хранить/получать заказы
- Сервис должен хранить все данные в кеше

## Задание №2

- Написать http server
- Server предоставляет API методы для взаимодействия с вашим сервисом
- Server Users слушает 80 порт
- Server Orders слушает 81 порт
- Для всех операций использовать method POST

## Задание №3

- Обернуть http server в docker container
