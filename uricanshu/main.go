package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:name/:age/:year/:month", func(ctx *gin.Context) {
		//获取路径参数
		name := ctx.Param("name")
		age := ctx.Param("age")
		year := ctx.Param("year")
		month := ctx.Param("month")
		ctx.JSON(http.StatusOK, gin.H{
			"name":  name,
			"age":   age,
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")
}
