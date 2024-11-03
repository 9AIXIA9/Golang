package test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 内容对齐
type s1 struct {
	a1 int8
	b1 string
	c1 int8
}

type s2 struct {
	a2 int8
	c2 int8
	b2 string
}

func TestStruct(t *testing.T) {
	a1 := s1{
		a1: 1,
		b1: "1",
		c1: 1,
	}
	a2 := s2{
		a2: 2,
		b2: "2",
		c2: 2,
	}
	fmt.Println(unsafe.Sizeof(a1), unsafe.Sizeof(a2))
}
