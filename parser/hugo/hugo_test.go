package hugo

import (
	"fmt"
	"testing"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

func TestParseImage(t *testing.T) {
	postItem := model.PostItem{
		Id:        "uffae",
		Title:     "",
		CreatedAt: "",
		Content:   "2016.11.18 - 2019.6.22\n![Image/1080/1920](https://media-image1.oss-cn-hangzhou.aliyuncs.com/studyroom_post_image/qmrmgu/3a2fc7490f0c59aa2dc3a01c5afc719b.c247c2af230d7409af8491a36eb9e52f.png)\n![Image/1080/1920](https://media-image1.oss-cn-hangzhou.aliyuncs.com/studyroom_post_image/qmrmgu/ebaed053961a3ca4599b47445931f9d2.a3eea3c46484c61883fcd13c30677398.png)",
	}
	h := HugoParser{}
	images := h.parseImage(&postItem)
	if len(images) != 2 {
		t.Error("图片解析失败！")
		return
	}
	for _, img := range images {
		fmt.Println(img.ContentUrl)
	}
	fmt.Println("=== content ===")
	fmt.Println(postItem.Content)
}

func TestParseContent(t *testing.T) {
	item := model.PostItem{
		Id:        "dsads",
		Title:     "",
		CreatedAt: "2019-06-22T15:00:11+0000",
		Content:   "大街上的徕卡肯德基卡萨鲁大\n师就来\n#\n看看 \n#\n萨达",
		Digest:    "大街上",
	}
	h := HugoParser{}
	result := h.parseContent(item)
	fmt.Println(result.SavePath)
	fmt.Println(result.BasicInfo)
}
