package main

import "fmt"

const defaultValueB = 100

// VerbOption 定义了一些配置项
type VerbOption struct {
	a string
	b int
	c bool
}

type OptionFunc func(*VerbOption)

func main() {
	o1 := NewVerbOption("only a")
	o2 := NewVerbOption("a and b", WithB(-100))
	fmt.Println(o1)
	fmt.Println(o2)
}

// NewVerbOption 选项模式使用 Verb
func NewVerbOption(a string, opts ...OptionFunc) *VerbOption {
	o := &VerbOption{a: a, b: defaultValueB}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithB 将 VerbOption 的 b 字段设置为指定值
func WithB(b int) OptionFunc {
	return func(o *VerbOption) {
		o.b = b
	}
}

// WithC 将 VerbOption 的 c 字段设置为指定值
func WithC(c bool) OptionFunc {
	return func(o *VerbOption) {
		o.c = c
	}
}
