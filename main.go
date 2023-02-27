package main

import (
	"database/sql"

	api "github.com/CRAZYKAYZY/web/api"
	"github.com/CRAZYKAYZY/web/cmd"
	"github.com/CRAZYKAYZY/web/db"
	sqlc "github.com/CRAZYKAYZY/web/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {

	go cmd.Execute()

	//call db func to connect
	db.ConnectDb()

	store := sqlc.NewStore(&sql.DB{})
	server := api.NewServer(store) //initialize server

	//start server
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
