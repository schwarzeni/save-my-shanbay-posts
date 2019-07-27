package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

// 获取文章的列表
func GetPostList(config model.AppConfig, pageNum int) (postList model.ReqPostList, err error) {
	var (
		client = http.Client{Timeout: 5 * time.Second} // TODO: 将其提出成配置文件
		req    *http.Request
		resp   *http.Response
		body   []byte
	)
	if req, err = http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://www.shanbay.com/api/v2/studyroom/users/%s/posts/?page=%d", config.UserID, pageNum),
		nil); err != nil {
		return
	}
	req.Header.Set("cookie", config.AuthToken)
	if resp, err = client.Do(req); err != nil {
		return
	}

	// 读取返回信息内容
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if err = json.Unmarshal(body, &postList); err != nil {
		return
	}

	return
}

// 获取文章的内容
func GetPostContent(postID string) (item model.ReqPostItem, err error) {
	var (
		client = http.Client{Timeout: 5 * time.Second} // TODO: 将其提出成配置文件
		req    *http.Request
		resp   *http.Response
		body   []byte
	)
	if req, err = http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://www.shanbay.com/api/v2/studyroom/posts/%s/", postID),
		nil); err != nil {
		return
	}
	if resp, err = client.Do(req); err != nil {
		return
	}

	// 读取返回信息内容
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if err = json.Unmarshal(body, &item); err != nil {
		return
	}

	return
}
