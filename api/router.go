package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Env struct {
	Db *sqlx.DB
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
