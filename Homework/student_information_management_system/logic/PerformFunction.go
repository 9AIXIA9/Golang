package logic

//进入菜单功能
import (
	"fmt"
	"os"
	"stuSys/models"
)

// PerformFunction 进入功能
func PerformFunction(choice int, students *[]models.Student) {
	switch choice {
	case 1:
		AddStudent(students)
		break
	case 2:
		EditStudent(students)
		break
	case 3:
		ShowStudents(students)
		break
	case 4:
		fmt.Println("期待您的下次使用！")
		os.Exit(0)
	}
}
