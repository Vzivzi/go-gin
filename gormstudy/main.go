package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type User struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&User{})

	u1 := User{Name: "kris", Gender: "男", Hobby: "篮球"}
	u2 := User{Name: "猪猪", Gender: "女", Hobby: "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// // 查询
	var u = new(User)
	// db.First(u)
	// fmt.Printf("%#v\n", u)

	var uu User
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// // 删除
	db.Delete(&uu)
}
