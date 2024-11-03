package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
	"strings"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:woshiXIJIA2005..@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db = sqlx.MustConnect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定。
	// 重新生成对应数据库的查询语句（如PostgreSQL 用 `$1`, `$2` bindvar）
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	IDs := make([]int, 5)
	IDs = append(IDs, 1, 2, 3, 4)
	Users := make([]User, 5)
	Users, _ = QueryByIDs(IDs)
	fmt.Println(Users)
	Users, _ = QueryAndOrderByIDs(IDs)
	fmt.Println(Users)

}
