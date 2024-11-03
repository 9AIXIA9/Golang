package main

import (
	"log"
	"sync"
	"time"
)

// 公交车站台
type stop struct {
	personNumber int //车上人数
	lock         *sync.Mutex
	t            *time.Ticker //周期循环
	c            chan bool    //判断车上是否满员
}

const (
	interval  = 20 * time.Second
	maxNumber = 30
)

func main() {
	s := &stop{
		lock: new(sync.Mutex),
		c:    make(chan bool, 1), // 使用缓冲通道避免阻塞
	}
	go s.newBus() // 在新的goroutine中运行
	// 模拟乘客上车
	for i := 0; i < 20; i++ {
		s.getOnFromMain(4)          // 从主goroutine调用，传递上车人数
		time.Sleep(1 * time.Second) // 模拟乘客陆续上车
	}
	time.Sleep(30 * time.Second) // 让程序运行一段时间以观察输出

}

// 乘客上车
func (s *stop) getOnFromMain(newPerson int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.personNumber += newPerson
	log.Printf("此时车上人数为%v", s.personNumber)
	if s.personNumber >= maxNumber {
		s.c <- true
		log.Print("人数已够，准备发车")
	}
}

// 新来一辆bus
func (s *stop) newBus() {
	s.t = time.NewTicker(interval)
	defer s.t.Stop() // 确保ticker在函数结束时被关闭
	for {
		select {
		case <-s.t.C:
			log.Print("时间到了，该发车了")
			s.launch()
		case <-s.c:
			log.Print("人数够了，该发车了")
			s.launch()
		}
	}
}

// 发车
func (s *stop) launch() {
	log.Print("此时发车了")
	s.lock.Lock()
	defer s.lock.Unlock()
	log.Printf("车上人数为%v", s.personNumber)
	s.personNumber = 0
}
