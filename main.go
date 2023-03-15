package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mikey247/go-bank/api"
	bankdb "github.com/mikey247/go-bank/db/sqlc"
	"github.com/mikey247/go-bank/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:5000"
// )

func main() {
	config,err:= util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn,err := sql.Open(config.DBdriver, config.DBsource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store:= bankdb.NewStore(conn)
	server:= api.NewServer(store)

	err= server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("cannot start server:",err)
	}
}

// func change (p *string){
//  fmt.Println(*p,p)
//  *p = "joey"
// }

// func main()  {
// 	var name = "ross"
// 	change(&name)
// 	fmt.Println(name)
// }