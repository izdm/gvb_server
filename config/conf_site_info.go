package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`
	BeiAn       string `yaml:"bei_an" json:"bei_an"`
	Title       string `yaml:"title" json:"title"`
	QQImage     string `yaml:"qq_image" json:"qq_image"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	SloganEn    string `yaml:"slogan_en" json:"sloganEn"`
	Web         string `yaml:"web" json:"web"`
	BilibiliUrl string `yaml:"bilibili_url" json:"bilibiliUrl"`
	GiteeUrl    string `yaml:"gitee_url" json:"giteeUrl"`
	GithubUrl   string `yaml:"github_url" json:"githubUrl"`
}
