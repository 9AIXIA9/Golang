package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//map的扩展使用 1

func main() {
	//按照固定顺序遍历map
	scoreMap := make(map[string]int, 100)
	//添加五十个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//按照从小到大的顺序去遍历scoreMap
	fmt.Println("---------------------")
	//1.取出所有的key存放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	//2.给切片中的key排序
	sort.Strings(keys)
	//目前keys是一个有序切片
	//3.按照排序输出key
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
