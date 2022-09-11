package api

import "github.com/gin-gonic/gin"

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok": true,
	})
}
