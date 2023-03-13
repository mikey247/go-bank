postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up
	# migrate -path db/migration -database "postgresql://root:container_password@localhost:5432/db_name?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down
	# migrate -path db/migration -database "postgresql://root:container_password@localhost:5432/db_name?sslmode=disable" -verbose down

sqlc:
	sqlc generate   

build:
	go build -v ./...

test:
	go test -v -cover ./...

server:
	go run main.go
