package logic

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func encrypt(oPassword string) (ePassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(oPassword), bcrypt.DefaultCost) //加密处理
	if err != nil {
		zap.L().Error("logic encrypt password failed", zap.Error(err))
		return "", err
	}
	ePassword = string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return ePassword, nil
}
