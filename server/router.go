package server

import (
	"app/controllers"
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

// embed load files only under current directory (cannot load files in parent directory)
//go:embed view/*.html
var f embed.FS

func NewRouter() *gin.Engine {
	router := gin.New()
	templ := template.Must(template.New("").ParseFS(f, "view/*.html"))
	router.SetHTMLTemplate(templ)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", controllers.IndexDisplayAction)
	router.GET("/collatz", controllers.InitialCollatzAction)
	router.POST("/calc_collatz", controllers.CollatzDisplayAction)
	router.GET("/ishealthy", controllers.IshealthyDisplayAction)
	return router
}
