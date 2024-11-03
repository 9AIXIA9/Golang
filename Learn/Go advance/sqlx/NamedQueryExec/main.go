package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:woshiXIJIA2005..@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db = sqlx.MustConnect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func insertUserDemo() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age":  28,
		})
	return
}
func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 1.使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("%#v\n", err)
			return
		}
	}(rows)
	for rows.Next() {
		var u user
		// StructScan is like sql.Rows.Scan, but scans a single Row into a single Struct.
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}

	// 2.使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("%#v\n", err)
			return
		}
	}(rows)
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	err = insertUserDemo()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	namedQuery()
}
