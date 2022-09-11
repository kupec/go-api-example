package main

import (
	"database/sql"
	"log"

	"example.com/api/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "local.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.NewRouter(api.Env{
		Db: db,
	})
	router.Run()
}
