// http返回和请求的全部数据格式
package model

// 请求返回的单个文件json数据结构
type PostItem struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
	Digest    string `json:"digest"`
}

// 请求返回的文章列表
type ReqPostList struct {
	StatusCode int    `json:"status_code"` // 请求状态
	Msg        string `json:"msg"`
	Data       struct {
		Objects []PostItem `json:"objects"`
	} `json:"data"` // 返回的文章列表信息
}

type ReqPostItem struct {
	StatusCode int      `json:"status_code"` // 请求状态
	Msg        string   `json:"msg"`
	Data       PostItem `json:"data"`
}
