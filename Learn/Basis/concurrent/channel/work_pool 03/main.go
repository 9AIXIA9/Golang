package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
	close(results)
}

// work pool
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启三个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	//发送五个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for ret := range results {
		fmt.Println(ret)
	}

}
