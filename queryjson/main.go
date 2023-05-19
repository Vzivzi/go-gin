package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//GET请求后面是querystring参数
	//key-value格式，多个用&连接
	//例如 /web?query=kris&age=19
	r.GET("/web", func(ctx *gin.Context) {
		//获取浏览器发送携带的query string参数
		name := ctx.Query("query") //获取到参数
		age := ctx.Query("age")
		//取不到就用默认值
		// name := ctx.DefaultQuery("query", "somebody")

		//布尔值版本
		// name, ok := ctx.GetQuery("query")
		// if !ok {
		// 	//取不到值返回布尔值false,ok一般用来接收布尔值
		// 	name = "somebody"
		// }
		ctx.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
