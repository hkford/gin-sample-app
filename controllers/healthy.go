package controllers

import (
	"github.com/gin-gonic/gin"
)

func IshealthyDisplayAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"IsHealthy": "healthy",
	})
}
