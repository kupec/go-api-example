package main

import (
	"log"
	"net/url"

	"example.com/api/api"
	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Api example starting")

	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.NewRouter(api.Env{
		Db: db,
	})
	router.Run(config.ApiBind)
}

func newConfig() (config api.Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println("Do not found .env file, using environment instead")
	}

	err = env.Parse(&config)
	return
}

func initDatabase(config api.Config) (db *sqlx.DB, err error) {
	dbURL, err := url.Parse(config.DatabaseURL)
	if err != nil {
		return
	}

	db, err = sqlx.Connect(dbURL.Scheme, config.DatabaseURL)
	if err != nil {
		return
	}

	return
}
