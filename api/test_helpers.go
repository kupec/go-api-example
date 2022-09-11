package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SetupRouter() (*sqlx.DB, *gin.Engine) {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec("CREATE TABLE users(id INT, name varchar(100), visits INT)")

	router := NewRouter(Env{
		Db: db,
	})
	gin.SetMode(gin.ReleaseMode)

	return db, router
}

func DoRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
