package post

import (
	"fmt"
	"testing"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

var config = model.AppConfig{
	UserID:       "qmrmgu",
	AuthToken:    "auth_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MjM4NjIzMDEsImV4cCI6MTU2NDkxNjYyMSwiZGV2aWNlIjoiIiwidXNlcm5hbWUiOiJzY2h3YXJ6ZW5pIiwiaXNfc3RhZmYiOmZhbHNlfQ.Z7ZfFXdwnf8J-Qx3N1Duynp3hxBaLKtHTObriZcPl6U",
	BlogPathRoot: ""}

func TestGetPostList(t *testing.T) {
	if fileList, err := GetPostList(config, 1); err != nil {
		t.Error(err)
	} else {
		if fileList.StatusCode != 0 {
			t.Error("状态码不为0,", fileList.Msg)
		} else {
			fmt.Println(fileList.Data)
		}
	}
}

func TestGetPostContent(t *testing.T) {
	if item, err := GetPostContent("mxacx"); err != nil {
		t.Error(err)
	} else {
		if item.StatusCode != 0 {
			t.Error("状态码不为0,", item.Msg)
		} else {
			fmt.Println(item.Data)
		}
	}
}
