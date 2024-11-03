package main

import (
	"fmt"
	"sync"
	"time"
)

var sy sync.WaitGroup

var mutex sync.Mutex

func main() {
	sy.Add(1)

	go func() {
		t := time.Now().Unix()
		testA()
		t1 := time.Now().Unix()
		fmt.Println("testA花费", t1-t)
		sy.Done()
	}()

	sy.Add(1)
	go func() {
		t := time.Now().Unix()
		testB()
		t1 := time.Now().Unix()
		fmt.Println("testB花费", t1-t)
		sy.Done()
	}()

	sy.Wait()
}

func testA() {
	var num = 0
	for i := 0; i < 1000; i++ {
		num++
		count()
	}
	fmt.Println("testA num为", num)
}

func testB() {
	num := 0
	//接受很多请求
	for i := 0; i < 10; i++ {
		sy.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				count()
				num++
			}
			sy.Done()
		}()
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("testB num为", num)
		}()
	}

}

func count() {
	count := 0
	for i := 0; i < 99999; i++ {
		count++
	}
}
