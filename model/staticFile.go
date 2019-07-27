package model

// 图片的信息
type ImageInfo struct {
	FromUrl    string // 来源的url，需要发送http请求来抓取
	SavePath   string // 保存的路径
	Name       string // 图片的名称
	ContentUrl string // 将其嵌入到静态博客中的图片路径
}

// 发布博客时的信息
type BlogPostInfo struct {
	BasicInfo PostItem // 基础信息
	SavePath  string   // 保存路径
}

// 需要保存文件的信息
type SaveFileInfo struct {
	Data        []byte // 文件内容
	FilePathStr string // 文件保存路径
}
