package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestPostCreateHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, PostCreateHandler)

	body := `{
    "title":"test",
    "content":"just a test"	,
    "community_id": "4"
} `
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//判断响应的内容是否按预期返回了需要登录的错误

	//方法1：判断响应内容中是否包含指定字符串
	assert.Contains(t, w.Body.String(), "需要登录")

	//方法2:将响应的内容反序列化到res 然后判断各字段与预期是否一致
	res := new(ResponseData)
	err := json.Unmarshal(w.Body.Bytes(), res)
	if err != nil {
		t.Fatalf("json unmarshal w.body failed,err:%v", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
