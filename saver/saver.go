package saver

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

// 保存文件
func Save(saveFileInfo model.SaveFileInfo) (err error) {
	if _, err = os.Stat(saveFileInfo.FilePathStr); os.IsNotExist(err) { // 当文件不存在时才进行保存操作
		dir := filepath.Dir(saveFileInfo.FilePathStr)
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return
		}

		if err = ioutil.WriteFile(saveFileInfo.FilePathStr, saveFileInfo.Data, 0777); err != nil {
			return
		}
	}
	return nil
}
