package main

import (
	"database/sql"
	"log"

	"example.com/api/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api.StartApi(api.Env{
		Db: db,
	})
}
