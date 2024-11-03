package mysql

import (
	"fmt"
	"life/internal/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Init 初始化mysql连接
func Init() (err error) {
	config := settings.Config.MysqlConf
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
		return fmt.Errorf("%w,connect mysql database failed,err:%w", errorMysqlInit, err)
	}
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	return
}

func Close() error {
	return db.Close()
}
