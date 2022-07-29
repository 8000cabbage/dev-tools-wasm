package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tool(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "html/tool.html", gin.H{
		"msg": "easy gin",
	})
}

func Json2Go(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool.html", gin.H{
		"method": "json2go",
		"notice": "提示",
		//"error":  "错误信息",
		"msg": "基于 https://github.com/m-zajac/json2go 包",
	})
}

func Sql2Gorm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tool.html", gin.H{
		"method": "sql2gorm",
		"notice": "提示",
		//"error":  "错误信息",
		"msg": "基于 https://github.com/cascax/sql2gorm 包",
	})
}

//func Yaml2Go(ctx *gin.Context) {
//	ctx.HTML(http.StatusOK, "tool.html", gin.H{
//		"method": "yaml2go",
//		"notice": "提示",
//		//"error":  "错误信息",
//		"msg": "基于yaml2go包",
//	})
//}
//func Toml2Go(ctx *gin.Context) {
//	ctx.HTML(http.StatusOK, "tool.html", gin.H{
//		"method": "toml2go",
//		"notice": "提示",
//		//"error":  "错误信息",
//		"msg": "基于toml2go包",
//	})
//}
