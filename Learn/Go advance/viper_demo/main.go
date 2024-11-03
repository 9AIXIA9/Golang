package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func main() {
	//设置默认值
	viper.SetDefault("fileDir", "./")
	//读取配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()           // 查找并读取配置文件
	if err != nil {                       // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//实时监控配置文件变化
	viper.WatchConfig()
	//当配置变化后调用一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		//当配置文件变化后调用的回调函数
		fmt.Println("config file changed:", e.Name)
	})
	r := gin.Default()
	r.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run()
}
