package main

import "fmt"

//数组

func main() {
	//1.定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "深圳", "广州"}
	fmt.Println(cityArray)

	//2.编译器推导数组的长度
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)

	//3.使用索引值方式初始化
	var langArray = [...]string{1: "golang", 2: "python", 8: "c++"}
	fmt.Println(langArray)

	//数组遍历
	//1.for遍历
	//for i := 0; i < len(cityArray); i++ {
	//	fmt.Println(cityArray[i])
	//}

	//2.for range遍历
	//有索引值
	//for index, value := range cityArray {
	//	//	fmt.Println(index, value)
	//	//}
	//	////使用_占据索引值
	//	//for _, value := range cityArray {
	//	//	fmt.Println(value)
	//	//}
	//数组是值类型
}
