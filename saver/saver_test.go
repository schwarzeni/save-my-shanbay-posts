package saver

import (
	"testing"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
)

func TestSave(t *testing.T) {
	if err := Save(model.SaveFileInfo{[]byte("112333"), "/Users/nizhenyang/Desktop/tmp/hugo-blog/static/images/zzz/test.txt"}); err != nil {
		t.Error(err)
	}
}
