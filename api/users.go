package api

import (
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func (env Env) listUsers(c *gin.Context) {
	users, err := models.GetUsers(env.Db)
	if err != nil {
		ResponseWithError(c, err)
		return
	}

	c.JSON(200, users)
}
