package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//匿名导入
)

// 定义一个全局对象db
var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

func initDB() (err error) {
	//DSN: Data Source Name

	dsn := "root:woshiXIJIA2005..@tcp(127.0.0.1:3306)/go_test"
	//此处不会检验账号密码是否正确
	db, err = sql.Open("mysql", dsn)

	//全局变量不要再次为其初始化:=
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("成功进入数据库")

}
