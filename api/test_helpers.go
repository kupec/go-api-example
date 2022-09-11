package api

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetupRouter() (*sql.DB, *gin.Engine) {
	db, err := sql.Open("sqlite3", "file:test.db?mode=memory")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE users(id INT, name varchar(100), visits INT)")
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter(Env{
		Db: db,
	})

	return db, router
}

func DoRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
