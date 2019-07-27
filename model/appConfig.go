package model

// 用户配置信息的配置
type AppConfig struct {
    UserID string `json:"user_id"` // 用户的ID
    AuthToken string `json:"auth_token"`
    BlogPathRoot string `json:"blog_path_root"` // 博客的根目录路径  不要以 '/' 结尾
}
