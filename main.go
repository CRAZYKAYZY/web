package main

import (
	//"database/sql"
	"log"

	api "github.com/CRAZYKAYZY/web/api"
	"github.com/CRAZYKAYZY/web/cmd"
	"github.com/CRAZYKAYZY/web/db"
	sqlc "github.com/CRAZYKAYZY/web/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {

	go cmd.Execute()

	//call db func to connect, to call a func from another package, ensure it returns a value.
	dbCon, err := db.ConnectDb()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	store := sqlc.NewStore(dbCon)
	server := api.NewServer(store) //initialize server

	//start server
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("error with server", err)
	}
}
