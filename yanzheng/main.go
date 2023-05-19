package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler func
func indexHandler(c *gin.Context) {
	fmt.Printf("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义中间件
func test(c *gin.Context) {
	fmt.Print("我来验证了")
	//计时
	start := time.Now()
	c.Next()
	cost :=
		time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("计时完成")
}

func yanzheng(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者验证工作
	return func(c *gin.Context) {
		if doCheck {
			fmt.Printf("恭喜您验证通过，即将跳转")
			c.Next()
		} else {
			fmt.Printf("登录失败，请重新登陆!")
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	//全局注册中间件

	r.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "shop",
		})
	})

	//只有这个用到test
	r.GET("./index", test, indexHandler)

	r.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "user",
		})
	})

	//group
	r.LoadHTMLFiles("./login.html", "./index.html")
	usersGroup := r.Group("users")
	{
		usersGroup.POST("/index", yanzheng(true), func(ctx *gin.Context) {
			username := ctx.DefaultPostForm("username", "未填写")
			password := ctx.PostForm("password")
			fmt.Printf("%v\n", password)
			ctx.HTML(200, "index.html", gin.H{
				"username": username,
			})
		})
		usersGroup.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(200, "login.html", nil)
		})

	}
	r.Run(":9090")
}
