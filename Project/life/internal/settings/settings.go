package settings

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Config      = new(AppConfig) // Config 全局变量 用来保存程序的所有配置信息
	errorConfig = errors.New("settings init failed")
)

// Init 初始化配置文件
func Init() error {
	//解决路径问题,相对路径切换目录后，无法正确找到文件
	var confFilePath string
	flag.StringVar(&confFilePath, "confFilePath", "", "配置文件路径")
	flag.Parse()

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w,get current work directory failed,error:%v", errorConfig, err)
	}

	// 如果提供的是相对路径，将其转换为绝对路径
	if !filepath.IsAbs(confFilePath) {
		confFilePath = filepath.Join(currentDir, confFilePath)
	}

	// 移除 可能存在的.exe 后缀
	confFilePath = strings.TrimSuffix(confFilePath, ".exe")

	// 规范化路径
	// 清理路径中的 .. 和 . 符号，删除多余的斜杠等
	// 例如："/home/user/../user/./settings.yaml" 会变成 "/home/user/settings.yaml"
	confFilePath = filepath.Clean(confFilePath)

	//检查是否传入配置文件
	if confFilePath == "" {
		return fmt.Errorf("%w,no settings file passed", errorConfig)
	}
	// 检查文件是否存在
	if _, err = os.Stat(confFilePath); os.IsNotExist(err) {
		return fmt.Errorf("%w,settings file not exist", errorConfig)
	}

	viper.SetConfigFile(confFilePath)
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		return fmt.Errorf("%w,read settings file failed,err:%w", errorConfig, err)
	}
	//把读取的配置信息反序列化到 Conf
	if err = viper.Unmarshal(Config); err != nil {
		return fmt.Errorf("%w,unmarshal settings file failed,err:%w", errorConfig, err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("settings file changed,try to reload...")
		if err := viper.Unmarshal(Config); err != nil {
			//viper.OnConfigChange()函数是默认保持最后一次加载成功的config
			//就算此次失败了，也不会造成程序的崩溃
			log.Printf("settings file reload failed,err: %v", err)
		} else {
			log.Println("settings file reload success")
		}
	})
	return nil
}
