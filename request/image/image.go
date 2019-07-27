package image

import (
	"io/ioutil"
	"net/http"
)

// 获取图片信息
func GetImage(imageUrl string) (result []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(imageUrl); err != nil {
		return
	}
	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
