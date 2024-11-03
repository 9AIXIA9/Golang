package main

import "fmt"

//切片

func main() {
	//切片实际上就是数组的包装

	//基于数组得到的切片
	a := [...]int{1, 2, 3, 4, 5, 6}
	b := a[1:3]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//基于切片的切片
	c := b[0:]
	//自动与结尾对齐
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	//make函数构造切片
	d := make([]int, 5, 10)
	//通过len函数求切片长度
	fmt.Println("D:", len(d))
	//通过cap函数求切片容量
	fmt.Println(cap(d))
	/*
		切片之间不能直接进行比较
		唯一能进行的只有与nil比较
	*/
	var x []int     //声明不赋值
	var y = []int{} //声明并初始化
	z := make([]int, 0)
	fmt.Println(x, len(x), cap(x))
	if x == nil {
		fmt.Println("x 是一个nil")
	}
	//x未赋值，未占据内存空间

	//切片的扩容
	x = append(x, 10)
	//此时增加的是一个元素
	fmt.Println(x)
	//多个元素扩容
	x = append(x, 1, 2, 3, 4, 5, 6)
	fmt.Println(x)
	//与另一切片进行扩容
	xxx := []int{11, 12, 13, 14}
	x = append(x, xxx...)
	fmt.Println(x)
	if x == nil {
		fmt.Println("x 是一个nil")
	} else {
		fmt.Println("现在x不是一个nil")
	}
	fmt.Println(y, len(y), cap(y))
	if y == nil {
		fmt.Println("y是一个nil")
	}
	//y进行了赋值，占据了内存空间
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("z 是一个nil")
	}
	//一般通过len函数来判断切片是否被赋初值

	//切片的赋值拷贝
	xx := make([]int, 3) //[0,0,0]
	yy := xx
	yy[1] = 100
	fmt.Println(xx)
	fmt.Println(yy)
	//说明指向同一个数组
	yy[2] = 200
	for index, value := range xx {
		fmt.Println(index, value)
	}
	for i := 0; i < len(xx); i++ {
		fmt.Println(xx[i])
	}
	//切片的复制(操作后并不指向同一数组)
	k := []int{1, 2, 3, 4, 5, 6}
	j := make([]int, 6, 6)
	copy(j, k)
	fmt.Println(j)
	fmt.Println(k)
	j[5] = 100
	fmt.Println(j)
	fmt.Println(k)
	//切片删除元素
	t := []string{"一", "二", "三", "四", "五"}
	t = append(t[0:2], t[3:]...)
	fmt.Println(t)
	/*
		一般来说通过此公式
		x = append(x[0：index],t[index + 1]...)
		index即为要删除的位置
	*/
	//slice的遍历
	for i, v := range k {
		fmt.Println(i, v)
	}
}
