package logic

//展示全体学生信息
import (
	"fmt"
	"stuSys/models"
)

func ShowStudents(students *[]models.Student) {
	fmt.Printf("%+v\n", *students)
}
