package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbCfg := settings.MysqlConfig{
		Host:               "127.0.0.1",
		User:               "root",
		Password:           "woshiXIJIA2005..",
		DatabaseName:       "bluebell",
		Port:               "3306",
		MaxOpenConnections: 20,
		MaxIdleConnections: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {

	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed,err:%v", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
