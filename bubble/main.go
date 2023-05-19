package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root1234@(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
func main() {

	//连接数据库
	//sql CREATE DATABASE BUBBLE；
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	//绑定模型
	DB.AutoMigrate(&Todo{})
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//add
		v1Group.POST("/todo", func(ctx *gin.Context) {
			var todo Todo
			ctx.BindJSON(&todo)
			//存入数据库
			//返回响应
			if err = DB.Create(&todo).Error; err != nil {
				ctx.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, todo)
			}
		})
		//look all
		v1Group.GET("/todo", func(ctx *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, todoList)
			}
		})
		//look one(id)
		v1Group.GET("todo/:id", func(ctx *gin.Context) {
			id, _ := ctx.Params.Get("id")
			var todo Todo
			if err = DB.Where("id=?", id).Find(&todo).Error; err != nil {
				ctx.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, todo)
			}
		})

		//update
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(200, gin.H{"error": "id不存在!"})
				return
			}

			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				ctx.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.BindJSON(&todo)
				DB.Save(&todo)
				ctx.JSON(200, todo)
			}
		})

		//delete
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(200, gin.H{"error": "id不存在!"})
				return
			}
			if err = DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				ctx.JSON(200, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, gin.H{id: "deleted"})
			}
		},
		)
	}
	r.Run(":9090")
}
