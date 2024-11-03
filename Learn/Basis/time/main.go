package main

import (
	"fmt"
	"time"
)

//time demo

func main() {
	//时间对象
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	//时间戳:从1970年1月1日到现在的一个秒数
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp1, timeStamp2)
	//将时间转换为具体的时间格式
	t := time.Unix(1715502269, 0)
	fmt.Println(t)
	//时间间隔
	time.Sleep(2 * time.Second)
	fmt.Println("over")
	//now + 1 hour
	fmt.Println(now)
	fmt.Println(now.Add(1 * time.Hour))
	//now - 1 hour
	fmt.Println(now.Sub(now))
	//定时器
	//for tep := range time.Tick(time.Second) {
	//	fmt.Println(tep)
	//}
	//时间格式化
	ret1 := now.Format("2006.01.02")
	ret2 := now.Format("2006-01-02.000")
	//2006 01 02 15 04 05
	// 年  月  日 时  分  秒
	fmt.Println(ret1)
	fmt.Println(ret2)
	//解析字符串类型的时间
	timeStr := "2019/08/0715:00:00"

	//1.拿到时区
	loc, err := time.LoadLocation("Asia/shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.根据时区去解析一个字符串格式的时间

	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj2)
}
