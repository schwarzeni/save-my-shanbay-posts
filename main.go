package main

import (
	"flag"

	"github.com/schwarzeni/save-my-shanbay-posts/core"
)

func main() {
	// 获取配置文件路径
	fpath := flag.String("p", "config.json", "config file")
	flag.Parse()

	core.Run(*fpath)
}
