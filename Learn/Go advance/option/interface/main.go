package main

//使用接口防止了对外暴露具体的配置结构体
import "fmt"

// verbOption 定义一个内部使用的配置项结构体
// 类型名称及字段的首字母小写（包内私有）
type verbOption struct {
	a string
	b int
	c bool
	// ...
}

// IOption 定义一个接口类型
type IOption interface {
	apply(*verbOption)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*verbOption)
}

func main() {
	DoSomething("only a")
	DoSomething("a and b", WithB(2121))
}

func (fo funcOption) apply(o *verbOption) {
	fo.f(o)
}

func newFuncOption(f func(*verbOption)) IOption {
	return &funcOption{
		f: f,
	}
}

// WithB 将b字段设置为指定值的函数
func WithB(b int) IOption {
	return newFuncOption(func(o *verbOption) {
		o.b = b
	})
}

// DoSomething 包对外提供的函数
func DoSomething(a string, opts ...IOption) {
	o := &verbOption{a: a}
	for _, opt := range opts {
		opt.apply(o)
	}
	// 在包内部基于o实现逻辑...
	fmt.Printf("o:%#v\n", o)
}
