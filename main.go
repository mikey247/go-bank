package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mikey247/go-bank/api"
	bankdb "github.com/mikey247/go-bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:5000"
)

func main() {
	conn,err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store:= bankdb.NewStore(conn)
	server:= api.NewServer(store)

	err= server.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start server:",err)
	}
}