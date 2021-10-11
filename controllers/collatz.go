package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitialCollatzAction(c *gin.Context) {
	c.HTML(200, "collatz.html", gin.H{})
}
