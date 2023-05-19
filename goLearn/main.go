package main // 声明 main 包，表明当前是一个可执行程序

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // ) // 导入内置 fmt 包

// // http server

// // func sayHello(w http.ResponseWriter, r *http.Request) {
// // 	b, _ := ioutil.ReadFile("./hello.txt")
// // 	_, _ = fmt.Fprintln(w, string(b))
// // }

// // func main() {
// // 	http.HandleFunc("/", sayHello)
// // 	err := http.ListenAndServe(":9090", nil)
// // 	if err != nil {
// // 		fmt.Printf("http server failed, err:%v\n", err)
// // 		return
// // 	}
// // }
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// 创建一个默认的路由引擎
// 	r := gin.Default()
// 	// GET：请求方式；/hello：请求的路径
// 	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
// 	r.GET("/goLn", func(c *gin.Context) {
// 		// c.JSON：返回JSON格式的数据
// 		c.JSON(200, gin.H{
// 			"message": "Hello world!",
// 		})
// 	})
// 	r.GET("/book", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "朱",
// 		})
// 	})
// 	r.POST("/bookk", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "post",
// 		})
// 	})
// 	r.PUT("/bookkk", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "PUT",
// 		})
// 	})
// 	r.DELETE("/bookkkk", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "DLEETE",
// 		})
// 	})
// 	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
// 	r.Run()
// }
