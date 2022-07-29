package main

import (
	"github.com/gin-gonic/gin"
	"go-tool/router"
	"net/http"
)

func main() {
	g := gin.Default()

	g.LoadHTMLGlob("template/*")
	g.StaticFS("/static", http.Dir("./static"))

	router.RegisterRoutes(g)

	g.Run(":8303")
}
