package logic

//添加新学生
import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"stuSys/models"
)

// AddStudent 添加学生
func AddStudent(students *[]models.Student) {
	var (
		id                 int
		idStr, name, class string
	)
	fmt.Println("请输入学生学号，姓名和班级(请用回车分隔)")
	_, err := fmt.Scanln(&idStr)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	id, err = strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	for _, r := range *students {
		if r.Id == id {
			err = errors.New("已存在该学号")
			log.Printf("%#v/n", err)
			return
		}
	}
	_, _ = fmt.Scanf("%s\n", &name)
	_, _ = fmt.Scanf("%s\n", &class)
	*students = append(*students, *NewStudent(id, name, class))
}
