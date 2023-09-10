package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Db     *sqlx.DB
	Config Config
}

func NewRouter(app App) *gin.Engine {
	r := gin.Default()

	r.GET("/health_check", healthCheck)
	r.GET("/users", app.listUsers)

	return r
}

func ResponseWithError(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}
