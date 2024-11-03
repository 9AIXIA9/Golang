package logic

//选择菜单功能
import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func FuncSelect() int {
	for {
		var choiceSrt string
		var choice int
		number, err := fmt.Scanln(&choiceSrt)
		//转字符串为数字
		choice, err = strconv.Atoi(choiceSrt)
		if err != nil {
			fmt.Println("输入错误")
			return 0
		}
		if err != nil || number != 1 {
			err = errors.New("输入数据格式有误！")
			log.Printf("%#v/n", err)
		} else if choice > 4 || choice < 1 {
			err = errors.New("输入数据大小有误！")
			log.Printf("%v/n", err)
		} else {
			return choice
		}
	}
}
