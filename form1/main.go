package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	//login post 把填写的信息发送出去
	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.DefaultPostForm("username", "未填写")
		password := ctx.PostForm("password")
		ctx.HTML(200, "index.html", gin.H{
			"Name":   username,
			"psword": password,
		})
	})
	r.Run(":9090")
}
