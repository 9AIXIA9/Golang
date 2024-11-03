package main
//任何gorm操作里面加上debug语句便可以看到具体实现的mysql语句
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//1.定义模型

type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小黑'"`
	Age  int64
}

func main() {

	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:woshiXIJIA2005..@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2.把模型和数据库里的表对应起来
	db.AutoMigrate(&User{})
	//3.创建
	u := User{
		Age: 32,
	}
	//所有零值不会保存到数据库但是会设置成默认值
	fmt.Println(db.NewRecord(&u)) //判断主值是否为空 true
	db.Create(&u)                 //在数据库中创建了一条 u的记录
	fmt.Println(db.NewRecord(&u)) //判断主值是否为空 false
	//目前已经弃用NewRecord，需要手动进行检查
	//要存入零值有两种方法
	//1.
	//u1 := User{Age: 12, Name: new(string)} //把Name类型改为指针
	//2.
	//u2 := User{Name: sql.NullString{String: "", Valid: true}, Age: 15}
	//可以用set方式实现插入
}
