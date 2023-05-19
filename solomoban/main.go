package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayGood(w http.ResponseWriter, r *http.Request) {
	//定义一个函数
	good := func(name string) (string, error) {
		return name + "真漂亮呀", nil
	}
	t := template.New("f.tmpl")
	// 告诉模版我现在多了一个自定义函数
	t.Funcs(template.FuncMap{
		"good": good,
	})
	_, err := t.ParseFiles("./f.tmpl")
	// t, err := template.New("f.tmpl").ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "猪猪"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "猪猪"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayGood)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
