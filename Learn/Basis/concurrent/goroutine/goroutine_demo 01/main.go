package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello 小王子")
	wg.Done()
	//通知wg已做完，计数器 - 1
}
func main() { //开启主goroutine去执行main函数

	wg.Add(1)  //计数牌 + 1
	go hello() //开启了一个独立的goroutine去执行hello这个函数
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait() //等所有小弟干完活才收兵
}
