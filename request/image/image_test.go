package image

import "testing"

func TestGetImage(t *testing.T) {
	if result, err := GetImage("https://media-image1.oss-cn-hangzhou.aliyuncs.com/studyroom_post_image/qmrmgu/3a2fc7490f0c59aa2dc3a01c5afc719b.c247c2af230d7409af8491a36eb9e52f.png"); err != nil {
		t.Error(err)
	} else if len(result) == 0 {
		t.Error("没有数据")
	}
}
