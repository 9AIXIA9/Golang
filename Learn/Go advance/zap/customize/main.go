package main

//自定义log配置，也可以用七米的
import (
	"encoding/json"
	"go.uber.org/zap"
)

func main() {
	rawJSON := []byte(`{
    "level":"debug",
    "encoding":"json",
    "outputPaths": ["stdout", "server.log"],
    "errorOutputPaths": ["stderr"],
    "initialFields":{"name":"dj"},
    "encoderConfig": {
      "messageKey": "message",
      "levelKey": "level",
      "levelEncoder": "lowercase"
    }
  }`)

	var cfg zap.Config                                    //进入细看
	if err := json.Unmarshal(rawJSON, &cfg); err != nil { //解析到config结构体里
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("server start work successfully!")
}
