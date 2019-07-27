package parser

import "github.com/schwarzeni/save-my-shanbay-posts/model"

// 将返回的数据解析成可供静态博客使用的文件
type Parser interface {
	Parse(postInfo model.PostItem) (blogPostInfo model.BlogPostInfo, imageUrl []model.ImageInfo)
	GetSavePath(postInfo model.PostItem) string
}
