package api

import (
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func (app App) listUsers(c *gin.Context) {
	users, err := models.GetUsers(app.Db)
	if err != nil {
		ResponseWithError(c, err)
		return
	}

	c.JSON(200, users)
}
