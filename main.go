package main

import (
	"log"

	"example.com/api/api"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "local.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.NewRouter(api.Env{
		Db: db,
	})
	router.Run()
}
