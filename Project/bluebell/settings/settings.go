package settings

import (
	"bluebell/myerrors"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Config 全局变量 用来保存程序的所有配置信息
var (
	Config = new(AppConfig)
)

// AppConfig 都必须用mapstructure
type AppConfig struct {
	App   *AppInfo `mapstructure:"app"`
	*Auth `mapstructure:"auth"`
	*Set  `mapstructure:"set"`

	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type Auth struct {
	TokenDuration int    `mapstructure:"token_duration"`
	TokenLocation string `mapstructure:"token_location"`
	TokenHeader   string `mapstructure:"token_header"`
	TokenSecret   string `mapstructure:"token_secret"`
}
type Set struct {
	PostVoteDuration int   `mapstructure:"post_vote_duration"`
	PostPerVoteScore int   `mapstructure:"post_per_vote_score"`
	PostDefaultPage  int64 `mapstructure:"post_default_page"`
	PostDefaultSize  int64 `mapstructure:"post_default_size"`
	CommunityDefault int64 `mapstructure:"community_default"`
}
type AppInfo struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      string `mapstructure:"port"`
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

func Init() (err error) {
	//解决路径问题,相对路径切换目录后，无法正确找到文件
	var confFilePath string
	flag.StringVar(&confFilePath, "confFilePath", "", "配置文件路径")
	flag.Parse()

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		return errors.Join(myerrors.ErrorGetCurrentDir, err)
	}

	// 如果提供的是相对路径，将其转换为绝对路径
	if !filepath.IsAbs(confFilePath) {
		confFilePath = filepath.Join(currentDir, confFilePath)
	}

	// 移除 可能存在的.exe 后缀
	confFilePath = strings.TrimSuffix(confFilePath, ".exe")

	// 规范化路径
	// 清理路径中的 .. 和 . 符号，删除多余的斜杠等
	// 例如："/home/user/../user/./config.yaml" 会变成 "/home/user/config.yaml"
	confFilePath = filepath.Clean(confFilePath)

	//检查是否传入配置文件
	if confFilePath == "" {
		return myerrors.ErrorConfPathNil
	}
	// 检查文件是否存在
	if _, err = os.Stat(confFilePath); os.IsNotExist(err) {
		return errors.Join(myerrors.ErrorConfPathNone)
	}
	viper.SetConfigFile(confFilePath)
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		return errors.Join(myerrors.ErrorViperRead, err)
	}
	//把读取的配置信息反序列化到 Conf
	if err = viper.Unmarshal(Config); err != nil {
		return errors.Join(myerrors.ErrorViperUnmarshal, err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err = viper.Unmarshal(Config); err != nil {
			fmt.Println(errors.Join(myerrors.ErrorViperUnmarshal, err))
		}
	})
	return
}
