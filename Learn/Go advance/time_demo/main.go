package main

//关于时间的一些基础操作
import (
	"fmt"
	"time"
)

func main() {
	timeDemo()
	formatDemo()
	computeDemo()
	tickerAndTimerDemo()

}

// 关于time.time类型的操作
func timeDemo() {
	currentTime := time.Now()
	fmt.Println("年:", currentTime.Year())
	fmt.Println("月:", currentTime.Month())
	fmt.Println("日:", currentTime.Day())
	fmt.Println("小时:", currentTime.Hour())
	fmt.Println("分钟:", currentTime.Minute())
	fmt.Println("秒:", currentTime.Second())
	fmt.Println("纳秒:", currentTime.Nanosecond())
	fmt.Println("星期几:", currentTime.Weekday())
	fmt.Println("Unix 时间戳（秒）:", currentTime.Unix())
	fmt.Println("Unix 时间戳（纳秒）:", currentTime.UnixNano())
}

// 时间格式化
func formatDemo() {
	timeStr := "2024-10-11 14:45:02"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析时间出错:", err)
	} else {
		fmt.Println("解析后的时间:", parsedTime)
	}
	formattedTime := parsedTime.Format("2006.01.02 15.04.05")
	//三种常见形式
	fmt.Println("格式化后的时间:", formattedTime)
	fmt.Println(parsedTime.Format("2006-01-02"))  // 2024-10-11 年月日
	fmt.Println(parsedTime.Format("15:04:05"))    // 14:45:02 时钟
	fmt.Println(parsedTime.Format("03:04:05 PM")) // 02:45:02 PM PM/AM

}

// 时间计算
func computeDemo() {
	//时间间隔计算
	startTime := time.Now()
	time.Sleep(2 * time.Second) // 模拟耗时操作
	elapsed := time.Since(startTime)
	fmt.Println("操作耗时:", elapsed)
	//时间加法
	currentTime := startTime

	nextWeek := currentTime.Add(7 * 24 * time.Hour)
	fmt.Println("一周后的时间:", nextWeek)

	yesterday := currentTime.Add(-24 * time.Hour)
	fmt.Println("昨天的时间:", yesterday)
	//时间减法
	pastTime := time.Date(2024, 10, 1, 9, 0, 0, 0, time.Local)
	duration := currentTime.Sub(pastTime)
	fmt.Println("时间差:", duration)
	//currentTime.After(sometime) //判断current time是否在sometime后面 , 若在则true
	//before同理
}

// 周期性操作
// 在使用timer和ticker时要注意并发安全
func tickerAndTimerDemo() {
	//延时操作
	fmt.Println("延时 3 秒执行")
	time.Sleep(3 * time.Second)
	fmt.Println("延时结束")
	//延时后触发操作
	fmt.Println("2 秒后执行")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("执行了某操作")
	}
	//周期性触发
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("每秒触发一次，当前时间:", t)
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker 已停止")

	//也可以这么使用
	fmt.Println("另一种周期性操作实现方法")
	fmt.Println("start")
	go func() {
		ticker2 := time.Tick(time.Second) //定义一个1秒间隔的定时器
		for i := range ticker2 {
			fmt.Println(i) //通道每收到一次信息就记录一次日志
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("end")

	//一次性定时器
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("等待 3 秒，再执行timer")
	<-timer.C
	fmt.Println("Timer 触发")

}
