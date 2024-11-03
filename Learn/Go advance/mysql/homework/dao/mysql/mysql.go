package mysql

import (
	"database/sql"
)

var db *sql.DB

func InitDB() (err error) {
	//DSN: Data Source Name
	dsn := "root:woshiXIJIA2005..@tcp(127.0.0.1:3306)/go_homework"
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
