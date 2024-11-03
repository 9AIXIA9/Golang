package mysql

import (
	"bluebell/myerrors"
	"bluebell/settings"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const driverName = "mysql"

var db *sqlx.DB

func Init(config *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect(driverName, dsn)
	if err != nil {
		zap.L().Error(myerrors.DBConnect.Error(), zap.Error(err))
		return
	}
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	return
}
func Close() {
	_ = db.Close()
}
