package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup

func main() { //开启主goroutine去执行main函数

	for i := 0; i < 10000; i++ {
		i := i
		wg.Add(1) //计数牌 + 1
		go func() {
			fmt.Println("hello", i)
			wg.Done()
		}() //开启了一个独立的goroutine去执行hello这个函数
	}

	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait() //等所有小弟干完活才收兵
}
