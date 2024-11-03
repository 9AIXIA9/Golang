package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	rawMessageDemo()
	url := "https://www.liwenzhou.com/posts/Go/json-tricks/"
	JSONEncodeDontEscapeHTML(URLInfo{url})
}

type sendMsg struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

// 处理不确定层级的json字符串
func rawMessageDemo() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`
	// 定义一个map，value类型为json.RawMessage，方便后续更灵活地处理
	var data map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("msg:%#v\n", msg)
	// msg:main.sendMsg{User:"q1mi", Msg:"永远不要高估自己"}
}

// URLInfo 一个包含URL字段的结构体
type URLInfo struct {
	URL string
	// ...
}

// JSONEncodeDontEscapeHTML json序列化时不转义 &, < 和 >
// & \u0026
// < \u003c
// > \u003e
func JSONEncodeDontEscapeHTML(data URLInfo) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json.Marshal(data) failed, err:%v\n", err)
	}
	fmt.Printf("json.Marshal(data) result:%s\n", b)

	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 告知encoder不转义
	if err := encoder.Encode(data); err != nil {
		fmt.Printf("encoder.Encode(data) failed, err:%v\n", err)
	}
	fmt.Printf("encoder.Encode(data) result:%s\n", buf.String())
}
