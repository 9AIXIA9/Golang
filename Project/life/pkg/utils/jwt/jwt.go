package jwt

//生成和解析token
import (
	"errors"
	"fmt"
	"life/internal/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	errorGenToken     = errors.New("generate token failed")
	errorInvalidToken = errors.New("token cant parse")
)

// MyClaims
// 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构
// 如果想要保存更多信息， 都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id" `
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userID int64, username string) (string, error) {
	var (
		mySecret      = []byte(settings.Config.TokenSecret)
		TokenDuration = time.Hour * time.Duration(settings.Config.TokenDuration)
	)
	//创建声明
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenDuration).Unix(),
			Issuer:    settings.Config.App.Name,
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		return "", fmt.Errorf("%w,err:%w", errorGenToken, err)
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	var mySecret = []byte(settings.Config.TokenSecret)
	//解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w,err:%w", errorInvalidToken, err)
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errorInvalidToken
}