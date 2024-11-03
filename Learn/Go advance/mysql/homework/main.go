package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn/dao/mysql"
	"learn/router"
	//匿名导入
)

// 定义一个全局对象db
var db *sql.DB

func main() {
	err := mysql.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("成功进入数据库")
	err = router.OpenWeb()
	if err != nil {
		fmt.Printf("open login web failed,err:%v\n", err)
		return
	}
}
