package seed

import (
	"fmt"
	"math/rand"
	"time"
)

// NewSeed 得到一个新种子
func NewSeed() *rand.Rand {
	// 创建一个新的随机数生成器实例，为其设置种子
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src)
}

// GetRandomCode 随机生成一个六位数验证码
func GetRandomCode() string {
	var randomCode string
	//时间戳作为种子的随机数
	s := NewSeed()
	for i := 0; i < 6; i++ {
		randomCode += fmt.Sprintf("%d", s.Intn(10))
	}
	return randomCode
}
