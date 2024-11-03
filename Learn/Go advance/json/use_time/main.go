package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Post struct {
	CreateTime time.Time `json:"create_time"`
}

type Post2 struct {
	CreateTime CustomTime `json:"create_time"`
}

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).UnixNano()

func main() {
	timeFieldDemo()
	timeFieldDemo2()
	customMethodDemo()
}

func timeFieldDemo() {
	p1 := Post{CreateTime: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	//内置的json包不识别我们常用的字符串时间格式，如2020-04-05 12:25:42
	var p2 Post
	if err = json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

func timeFieldDemo2() {
	p1 := Post2{CreateTime: CustomTime{time.Now()}}
	b, err := json.Marshal(p1)
	//json.Marshal()会检查CustomTime是否实现了json.marshal的接口
	//如果有就使用自定义接口
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post2
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

// UnmarshalJSON 把字节转化为time.time
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	fmt.Println("origin s:", s)
	s = strings.Trim(s, "\"") //去除引号
	fmt.Println("changed s:", s)

	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil //加引号
}

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Order struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedTime time.Time `json:"created_time"`
}

// MarshalJSON 为Order类型实现自定义的MarshalJSON方法
func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order // 定义与Order字段一致的新类型
	//注意他是type 没有自定义json序列化方法
	return json.Marshal(
		struct { //定义结构体,即解析 marshal格式
			CreatedTime string `json:"created_time"` //created_time序列化的归处
			*TempOrder         // 避免直接嵌套Order进入死循环
		}{
			CreatedTime: o.CreatedTime.Format(ctLayout), //把createTime转化为ctLayout格式
			TempOrder:   (*TempOrder)(o),                //
		})
}

// UnmarshalJSON 为Order类型实现自定义的UnmarshalJSON方法
func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order // 定义与Order字段一致的新类型
	ot := struct {
		CreatedTime string `json:"created_time"`
		*TempOrder         // 避免直接嵌套Order进入死循环
	}{
		TempOrder: (*TempOrder)(o),
	}
	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}
	var err error
	o.CreatedTime, err = time.Parse(ctLayout, ot.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}

// 自定义序列化方法
func customMethodDemo() {
	o1 := Order{
		ID:          123456,
		Title:       "《七米的Go学习笔记》",
		CreatedTime: time.Now(),
	}
	// 通过自定义的MarshalJSON方法实现struct -> json string
	b, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// 通过自定义的UnmarshalJSON方法实现json string -> struct
	jsonStr := `{"created_time":"2020-04-05 10:18:20","id":123456,"title":"《七米的Go学习笔记》"}`
	var o2 Order
	if err := json.Unmarshal([]byte(jsonStr), &o2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("o2:%#v\n", o2)
}
