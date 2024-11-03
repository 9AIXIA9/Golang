package main

//go demo1
import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql" //因为没有直接用到驱动，所以使用下划线

//UserInfo --> 数据库

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {

	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:woshiXIJIA2005..@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)

	}
	defer db.Close()
	//创建表  自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{
	//	ID:     1,
	//	Name:   "小王子",
	//	Gender: "男",
	//	Hobby:  "篮球",
	//}
	//db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //读取表格第一条数据于u
	fmt.Printf("%#v\n", u)
	//更新
	db.Model(&u).Update("Hobby", "羽毛球")
	db.Delete(&u)
}
