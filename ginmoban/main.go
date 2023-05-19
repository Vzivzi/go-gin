package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载静态文件，例如css样式给模板的
	r.Static("/static", "./statics")
	//添加自定义函数，比如相加一些过滤功能等等
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//模板解析
	r.LoadHTMLFiles("templates/index.tmpl", "users/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ //模版渲染
			"title": "<a href='liwenzhou.com'>我是一个网页~~</a>",
		})
	})
	r.GET("/index2", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "usersindex.tmpl", gin.H{ //模版渲染
			"title": "<a href='http://liwenzhou.com'>我是一个网页</a>",
		})
	})
	r.Run(":9090") //启动serve
}
