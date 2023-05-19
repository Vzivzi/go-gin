package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//一些数据
	r.GET("/index", func(ctx *gin.Context) {
		data := gin.H{"name": "kris", "sex": "man", "age": 20}
		ctx.JSON(200, data)
	},
	)
	//定义结构体然后传格式的数据
	type msg struct {
		Name string
		Age  int
		Sex  string
	}
	r.GET("/json", func(ctx *gin.Context) {
		data := msg{
			Name: "kris",
			Age:  19,
			Sex:  "man",
		}
		ctx.JSON(200, data)
	},
	)
	r.Run(":9090")
}
