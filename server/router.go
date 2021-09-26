package server

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("view/*.html")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", controllers.IndexDisplayAction)
	router.GET("/collatz", controllers.CollatzDisplayAction)
	router.GET("/ishealthy", controllers.IshealthyDisplayAction)
	return router
}
