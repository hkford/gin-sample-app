package server

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/ishealthy", controllers.IshealthyDisplayAction)
	return router
}
