package main

import "fmt"

//map的扩展使用 2

func main() {
	//值为map的切片
	var mapSlice = make([]map[string]int, 8)
	//仅完成了切片的初始化
	//还需要对map进行初始化
	mapSlice[0] = make(map[string]int, 8)
	//完成了map的初始化
	mapSlice[0]["彭于晏"] = 100
	fmt.Println(mapSlice)
	//值为切片的map
	var sliceMap = make(map[string][]int, 8)
	//只完成了map的初始化
	v, ok := sliceMap["中国"]
	fmt.Println(v)
	if ok {
		fmt.Println(sliceMap["中国"])
	} else {
		sliceMap["中国"] = make([]int, 4)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}
	//遍历sliceMap
	for value, i := range sliceMap {
		fmt.Println(i, value)
	}

}
