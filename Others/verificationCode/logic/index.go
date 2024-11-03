package logic

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"
	"verification/model"
	"verification/pkg/seed"
)

// CheckPhoneNumber 检测输入手机号的合理性
func CheckPhoneNumber() (string, error) {
	var input string
	fmt.Println("请输入手机号码")
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		log.Printf("fmt.Scanf failed,err:%v", err)
		return "", err
	}
	match, err := regexp.MatchString(`^1[3-9]\d{9}$`, input)
	if err != nil {
		log.Printf("regexp.MatchString failed,err:%v", err)
		return "", err
	}
	if !match {
		fmt.Println("手机号码格式错误")
		return "", errors.New(model.ErrForPhoneFormate)
	}
	return input, nil
}

// CheckIfHasExceedTimes 检测是否超过五次
func CheckIfHasExceedTimes(v int8) error {
	if v >= 5 {
		return errors.New(model.ErrForVerTimesHasExceed)
	}
	return nil
}

// CheckIfInDuration 检测是否在五分钟内
func CheckIfInDuration(i int64) error {
	if time.Now().UnixNano() < i {
		return errors.New(model.ErrForVerIntervalNotEnough)
	}
	return nil
}

// VerifyCode 认证逻辑
func VerifyCode(inputForCode string, inputForPhone string, p *map[string]model.PhoneDetail) bool {
	if inputForCode == (*p)[inputForPhone].Code {
		(*p)[inputForPhone] = model.PhoneDetail{
			ValidDuration: 0,
			Interval:      0,
			Times:         (*p)[inputForPhone].Times,
			Code:          seed.GetRandomCode(), // 重新生成验证码 防止重复使用
		}
		return true
	}
	return false
}
