package main

import (
	"stuSys/logic"
	"stuSys/models"
)

//学员信息管理系统

//1.添加学员信息
//2.编辑学员信息
//3.展示全部学员信息

func main() {
	students := make([]models.Student, 0, 10)

	for {
		//1.打印系统菜单
		logic.ShowMenu()

		//2.功能选项
		choice := logic.FuncSelect()

		//3.执行功能
		logic.PerformFunction(choice, &students)
	}
}
