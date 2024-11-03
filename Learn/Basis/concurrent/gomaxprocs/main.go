package main

import (
	"fmt"
	"runtime"
	"sync"
)

//gomaxprocs demo
var wg sync.WaitGroup

func A() {
	for i := 0; i < 100; i++ {
		fmt.Println("A", i)
	}
	wg.Done()
}
func B() {
	for i := 0; i < 100; i++ {
		fmt.Println("B", i)
	}
	wg.Done()
}
func main() {
	runtime.GOMAXPROCS(1) //只占用一个CPU核心 先做完一项工作在做另外
	//占用一个以上时同时进行
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
