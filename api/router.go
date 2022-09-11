package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Env struct {
	Db *sql.DB
}

func StartApi(env Env) {
	r := gin.Default()

	r.GET("/health_check", healthCheck)
	r.GET("/users", env.listUsers)

	r.Run()
}

func ResponseWithError(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}
