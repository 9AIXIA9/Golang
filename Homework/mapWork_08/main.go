package main

import (
	"fmt"
	"strings"
)

//map(映射)
//统计一个字符串中每个单词出现的次数
//"how do you do"中每个单词出现的次数

func main() {
	x := make(map[string]int, 50)
	str := "how do you do"
	y := strings.Split(str, " ")
	for i := 0; i < len(y); i++ {
		x[y[i]]++
	}
	for v, i := range x {
		fmt.Println(v, "的次数为", i)
	}
}
