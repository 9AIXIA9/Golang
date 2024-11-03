package init

import "fmt"

// init函数在导包时自动执行
// init函数既没参数也没返回值
// 全局变量 > init > main
// init多用来进行初始化操作
func init() {
	fmt.Println("我是init函数")
}
