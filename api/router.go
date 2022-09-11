package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Env struct {
	Db *sql.DB
}

func NewRouter(env Env) *gin.Engine {
	r := gin.Default()

	r.GET("/health_check", healthCheck)
	r.GET("/users", env.listUsers)

	return r
}

func ResponseWithError(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}
