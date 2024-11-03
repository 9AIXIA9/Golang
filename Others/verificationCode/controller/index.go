package controller

import (
	"fmt"
	"log"
	"time"
	"verification/logic"
	"verification/model"
	"verification/pkg/deposeMsg"
	"verification/pkg/seed"
)

// GetRandomCode 随机生成一个六位数验证码 验证码有效时间五分钟
func GetRandomCode(p *map[string]model.PhoneDetail, number string) {
	//得到数据
	value, exit := (*p)[number]
	//不存在就设为一
	if !exit {
		(*p)[number] = model.PhoneDetail{
			ValidDuration: (*p)[number].ValidDuration,
			Interval:      (*p)[number].Interval,
			Times:         0,
			Code:          (*p)[number].Code,
		}
	}
	//判断是否超过五次
	if err := logic.CheckIfHasExceedTimes(value.Times); err != nil {
		fmt.Printf("CheckIfHasExceedTimes failed,err:%v\n", err)
		return
	}
	//判断是否在五分钟内
	if err := logic.CheckIfInDuration(value.Interval); err != nil {
		fmt.Printf("CheckIfInDuration failed,err:%v\n", err)
		return
	}
	randomCode := seed.GetRandomCode()

	(*p)[number] = model.PhoneDetail{
		ValidDuration: time.Now().UnixNano() + 5*60*1e9,
		Interval:      time.Now().UnixNano() + 5*60*1e9,
		Times:         (*p)[number].Times + 1,
		Code:          randomCode,
	}
	fmt.Println("验证码为:", randomCode)
	_ = deposeMsg.Save(p)
}

// VerifyCode 认证逻辑
func VerifyCode(p *map[string]model.PhoneDetail) (bool, error) {
	var inputForPhone string
	var inputForCode string
	fmt.Println("请输入手机号")
	if _, err := fmt.Scanf("%s\n", &inputForPhone); err != nil {
		log.Printf("fmt.Scanf failed,err:%v", err)
		return false, err
	}
	value, ok := (*p)[inputForPhone]
	//检查是否过期
	if !ok || time.Now().UnixNano() > value.ValidDuration {
		fmt.Println(model.ErrorForCodeHasExpired)
		return false, nil
	}

	fmt.Println("请输入验证码")
	if _, err := fmt.Scanf("%s\n", &inputForCode); err != nil {
		log.Printf("fmt.Scanf failed,err:%v", err)
		return false, nil
	}
	if logic.VerifyCode(inputForCode, inputForPhone, p) {
		log.Println("验证成功")
		return false, nil
	}
	log.Println("验证码错误")
	return false, nil
}
