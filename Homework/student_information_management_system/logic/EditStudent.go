package logic

//编辑学生信息
import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"stuSys/models"
)

// EditStudent 编辑学生信息
func EditStudent(students *[]models.Student) {
	var (
		idStr string
		name  string
		class string
		ok    bool
	)

	fmt.Println("请输入修改的学生学号，姓名和班级(请用回车分隔)")
	//字符转数字
	_, err := fmt.Scanln(&idStr)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	id, err := strconv.Atoi(idStr)
	_, _ = fmt.Scanf("%s\n", &name)
	_, _ = fmt.Scanf("%s\n", &class)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	for i, v := range *students {
		if v.Id == id {
			v.Name = name
			v.Class = class
			ok = true
			(*students)[i] = v
			break
		}
	}
	if !ok {
		err = errors.New("系统无此学号！")
		log.Printf("%#v/n", err)
	}
}
