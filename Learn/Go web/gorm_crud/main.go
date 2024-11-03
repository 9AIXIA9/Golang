package main

//gorm中的查询
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1.定义模型

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// MySQL DSN（数据源名称）
	dsn := "root:woshiXIJIA2005..@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	// 2.打开 MySQL 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 连接成功
	fmt.Println("Database connection established")

	// 把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	//db.Create(&User{Name: "John", Age: 25})
	//db.Create(&User{Name: "Kobe", Age: 18})
	// 4.查询
	//单体查询
	var user User
	db.First(&user, 1) // 查询ID为1的用户(仅当主键为int类型时可以使用)
	// First代表匹配条件的第一条记录 Last代表匹配条件的最后一条记录 （就是升降序的差异）
	fmt.Println(user)
	//查询所有值
	var users []User
	db.Find(&users) //所有值
	fmt.Printf("%#v\n", users)
	//条件查询 Where(用法很多)
	//GORM只通过非零字段查询,零值如 0和false等不会构建查询条件
	db.Where("name = ? AND age >= ?", "John", 18).Find(&users)
	fmt.Println(users)
	//Not条件与Where类似且相反(排除某些条件符合的)
	//Or条件与Where类似 满足一个条件即可
	//内联条件(把条件写入语句内部)
	db.First(&users, "name = ?", "Kobe")
	fmt.Println(users)
}
