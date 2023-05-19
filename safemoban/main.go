package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("show.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./show.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	str1 := "<a href='https://www.baidu.com/'> 我是一个a标签</a>"
	str2 := "<script>alert('我是弹出来的信息哦')</script>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
