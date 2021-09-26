package server

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(IsLocal bool) *gin.Engine {
	router := gin.New()
	// If you run `go test ./test` then load html files from ../view/*.html
	if IsLocal {
		router.LoadHTMLGlob("../view/*.html")
	} else {
		router.LoadHTMLGlob("view/*.html")
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", controllers.IndexDisplayAction)
	router.GET("/collatz", controllers.CollatzDisplayAction)
	router.GET("/ishealthy", controllers.IshealthyDisplayAction)
	return router
}
