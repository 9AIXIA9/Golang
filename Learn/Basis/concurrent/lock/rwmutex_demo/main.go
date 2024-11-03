package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
var (
	x      int64
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func read() {
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg.Done()

}
func write() {
	rwLock.Lock()
	//lock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	//lock.Unlock()
	rwLock.Unlock()
	wg.Done()

}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
