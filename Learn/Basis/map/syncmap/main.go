package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map
// 原生map不支持并发
var (
	wg sync.WaitGroup
	//m  = make(map[int]int)
	m2 = sync.Map{}
)

//func get(key int) int {
//	return m[key]
//
//}
//func set(key int, value int) {
//	m[key] = value
//}

//	func main() {
//		for i := 0; i < 20; i++ {
//			wg.Add(1)
//			go func(i int) {
//				set(i, i+100)
//				fmt.Printf("key:%v value:%v", i, get(i))
//				wg.Done()
//			}(i)
//		}
//		wg.Wait()
//	}
func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
