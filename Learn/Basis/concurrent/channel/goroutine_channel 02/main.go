package main

import "fmt"

/*
两个goroutine,两个channel
1,生成0~100的数字发送到ch1
2,从ch1中取出数据计算它的平方，把结果发送到ch2中
*/

// 生成0 ~ 100的数字发送到ch1
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	//关闭后仍可以读取变量
}

// 从ch1中取出数据计算，然后储存到ch2
func f2(ch1 <-chan int, ch2 chan<- int) { //chan<-单向通道（只能被发送） <-chan（只能被接收）
	//从通道中取值的方式 1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)

	//从通道中取值的方式 2
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
