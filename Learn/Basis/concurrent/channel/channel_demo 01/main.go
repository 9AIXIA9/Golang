package main

import "fmt"

// channel demo
func main() {
	//var ch1 chan int//引用类型需要进行初始化才能进行使用
	//ch1 = make(chan int,1)
	//ch1 := make(chan int)    //无缓冲区通道  又称为同步通道
	ch1 := make(chan int, 1) //带缓冲区通道 后面的数字是缓冲区大小
	ch1 <- 10                //发送值
	x := <-ch1
	fmt.Println(x)
	//len(ch1) //取通道中元素数量
	//cap(ch1) //拿到通道容量
	close(ch1)
}
