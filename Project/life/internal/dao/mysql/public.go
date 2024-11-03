package mysql

//mysql公用词语
import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

// 数据库名称
var (
	db *sqlx.DB
)

// 常量
const (
	driverName   = "mysql"
	userInterval = 5 * time.Second
	capInterval  = time.Minute
)

// 错误
var (
	errorMysqlInit = errors.New("mysql init failed")
	errorOperate   = errors.New("operate mysql failed")
	ErrorDataNil   = errors.New("mysql not exist this data")
)
