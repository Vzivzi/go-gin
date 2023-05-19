package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// main.go

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	u1 := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	//另一种数据方法不用首字母大写
	m1 := map[string]interface{}{
		"name":   "kris",
		"gender": "man",
		"age":    19,
	}
	// 数组
	hobby := []string{
		"篮球",
		"羽毛球",
		"橄榄球",
	}
	// tmpl.Execute(w, user)
	tmpl.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobby,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
