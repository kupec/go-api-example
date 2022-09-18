package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func SetupRouter() (*sqlx.DB, *gin.Engine) {
	db := initDatabase()

	router := NewRouter(Env{
		Db: db,
	})
	gin.SetMode(gin.ReleaseMode)

	return db, router
}

func initDatabase() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")

	instance, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../models/migrations", "sqlite3", instance)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		panic(err)
	}

	return db
}

func DoRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
