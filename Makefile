postgres:
    docker run --name containerName -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres12 dropdb simple_bank

migrateup:
    migrate -path db/migration -database "postgresql://root:container_password@localhost:5432/db_name?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql://root:container_password@localhost:5432/db_name?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb