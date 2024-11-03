package main

import "fmt"

//select

func main() {
	ch1 := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:
		default:
			fmt.Println("啥都不干")
		}
	}
	fmt.Println("---------------------")
	ch2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { //每次从满足条件的随机挑选一个
		case x := <-ch2:
			fmt.Println(x)
		case ch2 <- i:
		default:
			fmt.Println("啥都不干")
		}
	}
}
