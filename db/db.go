package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() {
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	var (
		DB  *sql.DB
		err error
	)

	DB, err = sql.Open("postgres", pgConnString)
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	err = DB.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	return
}
