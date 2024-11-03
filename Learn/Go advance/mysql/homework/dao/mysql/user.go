package mysql

import (
	"fmt"
	"learn/model"
)

func Login(u model.User) bool {
	sqlStr := "select password from user where phone= ?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	var userInDB model.User
	err := db.QueryRow(sqlStr, u.Phone).Scan(&userInDB.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	if userInDB.Password != u.Password {
		return false
	}
	fmt.Println("成功登录!")
	return true
}

// CheckExist 用于检测用户是否存在
func CheckExist(u model.User) bool {
	sqlStr := "select id, phone from user where phone = ?"
	var userToUse model.User
	db.QueryRow(sqlStr, u.Phone).Scan(userToUse.Id, userToUse.Phone)
	if userToUse.Phone != 0 {
		return false
	}
	return true
}

// UpLoad 用于上传用户数据到数据库
func UpLoad(u model.User) {
	sqlStr := "insert into user(name, phone,password) values (?,?,?)"
	ret, err := db.Exec(sqlStr, u.Name, u.Phone, u.Password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	u.Id, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("成功上传数据库,the id is %d.\n", u.Id)
}
