package main

//使用手机号码验证系统
/*
问题；
验证码时间差，数据整合，信息保存
*/

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

const (
	DIGIT       = 1000000
	MAXTRYNUM   = 5
	MaxInterval = 300
)

func main() {
	usedNumber := make(map[uint]uint8, 5)
	nextTime := make(map[uint]time.Time, 5)
	var phoneNumber uint
	/*
		phoneNumber用于存储用户输入账号
		usedNumber用于存储每个号码的验证码发送次数
	*/
	//读入用户电话号码
	fmt.Println("请输入您的电话号码：")
	num, err := fmt.Scanf("%d", &phoneNumber) //415421
	//判断读入电话号码是否出错
	if err != nil || num == 0 {
		err = errors.New("电话号码格式有误，本系统只支持非负整数")
		fmt.Println(err)
		os.Exit(-1)
	}
	//判断此电话号码的验证码的使用次数是否上限和上次获取验证码的时间是否过短
	for {

		now := time.Now()
		if usedNumber[phoneNumber] > MAXTRYNUM {
			err = errors.New("此电话号码一天内已使用验证码超过最大次数")
			log.Printf("%v", err)
			return
		} else if nextTime[phoneNumber].Sub(now).Seconds() <= 0 {
			err = errors.New("不可多次获取验证码，请五分钟后重试！")
			fmt.Println(err)
		} else {
			//增加使用次数
			usedNumber[phoneNumber]++
			nextTime[phoneNumber] = time.Now().Add(time.Second * MaxInterval)
			fmt.Println("验证码为：")
			//生成随机数验证码
			myRand, _ := rand.Int(rand.Reader, big.NewInt(DIGIT))
			fmt.Printf("%6v\n", myRand)
		}
	}
}
