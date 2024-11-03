package settings

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Config 全局变量 用来保存程序的所有配置信息
var Config = new(AppConfig)

//都必须用mapstructure

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         string `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host               string `mapstructure:"host"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DatabaseName       string `mapstructure:"database_name"`
	Port               string `mapstructure:"port"`
	MaxOpenConnections int    `mapstructure:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	//viper.SetConfigFile("config.yaml")//直接指定配置文件
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("json")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	// (专用于从远程获取配置信息时指定配置文件类型)
	//本地时不会生效
	//viper.AddConfigPath("./settings") // 查找配置文件所在的路径
	if filePath == "" {
		return fmt.Errorf("配置文件路径为空")
	}
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在：%s", filePath)
	}
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig failed,err:%v\n", err)
		return
	}
	//把读取的配置信息反序列化到 Conf
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v", err)
		}
	})
	return
}
