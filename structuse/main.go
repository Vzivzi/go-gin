package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 参数绑定 大量数据可用 下面是一个数据的例子
type UserInfo struct {
	// 	username string
	// 	password string
	Username string `form:"username"`
	Password string `form:"password"`
	//必须大写字母，如果用绑定数据的方式
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	r.POST("/form", func(ctx *gin.Context) {
		// username := ctx.Query("username")
		// password := ctx.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo //声明一个变量u，利用上面的结构体

		//带一个指针才能修改数据不然错误
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(200, gin.H{
				"msg":      "OK",
				"username": u.Username,
				"password": u.Password,
			})
		}

		// fmt.Printf("%#v\n", u)
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"msg": "ok",
		// })
	})

	r.Run(":9090")
}
