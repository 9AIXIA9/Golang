package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	secret = "我的盐"
)

// encryptPassword 对用户密码进行加密
func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}

func main() {
	oPassword := "123456"
	fmt.Println(oPassword)
	fmt.Println(encryptPassword(oPassword))
}
