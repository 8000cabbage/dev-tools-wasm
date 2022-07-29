package router

import (
	"github.com/gin-gonic/gin"
	"go-tool/controller"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.GET("/other", func(c *gin.Context) {
		c.String(http.StatusOK, "other")
	})
	router.GET("/go_tool", func(c *gin.Context) {
		c.String(http.StatusOK, "go_tool")
	})

	router.GET("/go_tool/json2go", controller.Json2Go)
	router.GET("/go_tool/sql2gorm", controller.Sql2Gorm)
	//router.GET("/go_tool/yaml2go", controller.Yaml2Go)
	//router.GET("/go_tool/toml2go", controller.Toml2Go)
}
