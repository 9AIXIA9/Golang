package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 定义模型
//表名默认为结构体复数

type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` //零值类型
	Birthday     *time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`        // 设置字段大小255
	MemberNumber string `gorm:"unique;not null"` //设置会员号唯一且不为空
	Num          int    `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string `gorm:"index:addr"`      //给Address字段创建名为addr的索引
	IgnoreMe     int    `gorm:"-"`               //忽略本字段
}

//// 修改默认表名 (唯一默认指定表名)
//func (User) TableName() string {
//	return "小王子们"
//
//}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "前缀" + defaultTableName
	} //给默认表名加上前缀
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:woshiXIJIA2005..@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) //禁用复数
	db.AutoMigrate(&User{})
	db.Table("玫瑰").CreateTable(&User{}) //用user创建一个叫玫瑰的表

}
