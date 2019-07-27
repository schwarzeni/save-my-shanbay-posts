// 专门将内容解析成hugo形式的parser
package hugo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

type HugoParser struct {
	BlogRoot string
}

func (hp *HugoParser) Parse(postInfo model.PostItem) (blogPostInfo model.BlogPostInfo, imageUrl []model.ImageInfo) {
	imageUrl = hp.parseImage(&postInfo)
	blogPostInfo = hp.parseContent(postInfo)
	return
}

func (hp *HugoParser) GetSavePath(item model.PostItem) string {
	return fmt.Sprintf("%s/content/posts/%s_%s_%s.md", hp.BlogRoot, item.CreatedAt, item.Id, item.Title)
}

// 生成博客内容
func (hp *HugoParser) parseContent(item model.PostItem) (blogPostInfo model.BlogPostInfo) {
	// 修改日期格式
	item.CreatedAt = strings.Split(item.CreatedAt, "+")[0]

	// 将所有的 "\n#\n" 转变为 "\n\n"
	compile := regexp.MustCompile(`\n#\n`)
	item.Content = compile.ReplaceAllString(item.Content, "\n\n")

	// 添加首页概览
	// TODO 考虑是否添加digest，暂时先去掉，因为hugo可以为我们默认添加
	//if len(strings.Trim(item.Digest, " ")) > 0 {
	//	idx := strings.Index(item.Content, item.Digest) + len(item.Digest)
	//	// TODO: 可能会 out-out-bound
	//	if len(item.Content) >= idx {
	//		item.Content = item.Content[:idx] + "<!--more-->" + item.Content[idx:]
	//	}
	//}

	// 设置保存路径
	blogPostInfo.SavePath = hp.GetSavePath(item)

	// 如果标题不存在则添加标题
	if len(item.Title) == 0 {
		item.Title = strings.Split(item.CreatedAt, "T")[0]
	}

	// 生成模板
	item.Content = fmt.Sprintf(`
---
title: "%s"
date: %s
---

%s

`, item.Title, item.CreatedAt, item.Content)

	blogPostInfo.BasicInfo = item
	return
}

// 将图片抽取出来
func (hp *HugoParser) parseImage(item *model.PostItem) (images []model.ImageInfo) {
	// TODO 找出 findAndReplace的方法
	// ![dsasa](/image/url/dsdsa.png)
	compile := regexp.MustCompile(`\!\[[^\]]*\]\(([^\)]*)\)`)
	regexResult := compile.FindAllStringSubmatch(item.Content, -1)
	for _, match := range regexResult {
		fileName := filepath.Base(match[1])
		contentUrl := fmt.Sprintf("/images/%s/%s", item.Id, fileName)
		item.Content = strings.Replace(item.Content, match[1], contentUrl, -1)
		images = append(images, model.ImageInfo{
			FromUrl:    match[1],
			SavePath:   fmt.Sprintf("%s/static%s", hp.BlogRoot, contentUrl),
			Name:       fileName,
			ContentUrl: contentUrl,
		})
	}
	return
}
