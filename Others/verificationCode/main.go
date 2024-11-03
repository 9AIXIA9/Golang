package main

import (
	"fmt"
	"log"
	"verification/controller"
	"verification/logic"
	"verification/model"
	"verification/pkg/deposeMsg"
)

func main() {
	//用map来储存对应的手机号和相关信息
	var detailForPhone = make(map[string]model.PhoneDetail)
	var err error
	detailForPhone, err = deposeMsg.Get()
	if err != nil {
		log.Println("deposeMsg.Get failed,err:", err)
		return
	}
	fmt.Println("欢迎来到验证码系统\n请按照提示输入")
	for {
		fmt.Println("\t1.输入手机号码获取验证码\n\t2.输入手机号码并进行验证")
		var input int
		_, _ = fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			// 1.检测手机号合理性
			phone, err := logic.CheckPhoneNumber()
			if err != nil {
				if err.Error() == model.ErrForPhoneFormate {
					continue
				} else {
					return
				}
			}
			//	2.随机生成一个六位数验证码 验证码有效时间五分钟
			controller.GetRandomCode(&detailForPhone, phone)
			break
		case 2:
			// 3.认证逻辑
			ok, err := controller.VerifyCode(&detailForPhone)
			if err != nil {
				return
			}
			if ok {
				fmt.Println("验证成功")
				return
			}
			break
		default:
			fmt.Println("输入错误")
			break
		}
	}
}
