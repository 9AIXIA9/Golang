package logic

//逻辑错误做细分，服务出错化笼统
import "errors"

var (
	ErrorUserExist     = errors.New("user have existed")
	ErrorUserNotExist  = errors.New("user not exist")
	ErrorWrongPassword = errors.New("wrong username or password")
	ErrorTooFrequent   = errors.New("frequency is too high")
	ErrorCapExist      = errors.New("capability have existed")
	ErrorCapNotExist   = errors.New("capability not exist")
	ErrorUserNoneCap   = errors.New("user have not set capability")
)
