# glossary-service

## restful-service
Реализован сервис, который реализует операции к
глоссарию употребляемых терминов,
связанных с производительностью веб-приложений, использующих WebAssembly:

1. Получение списка всех терминов.
2. Получение информации о конкретном термине по ключевому слову.
3. Добавление нового термина с описанием.
4. Обновление существующего термина.
5. Удаление термина из глоссария.

Сервис реализован с использованием технологий:
- Go 1.23,
- Podman,
- Echo,
- Swagger,
- PostgreSQL.

## Генерация документации swagger

[Документация в формате swagger](docs/swagger.yaml).

```shell
swag init -g cmd/service/main.go
```

## Запуск и работа с API через swagger

```shell
podman compose up
```

```
http://localhost:8080/swagger/
```

## gRPC

В этот же репозиторий добавил реализацию gRpc сервиса и клиента для него.
Не изменяя реализации работы с базой данных, добавил еще один сервис на grpc.

```
  grpc-service:
    build:
      context: .
      dockerfile: deployments/grpc.Dockerfile
    environment:
      DATABASE_URL: postgres://user:password@db:5432/database?sslmode=disable
    ports:
      - "50051:50051"
    depends_on:
      - db
```
