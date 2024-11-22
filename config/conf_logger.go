package config

type Logger struct {
	Level        string `yaml:"level"`          // 日志级别，例如 info、debug 等
	Prefix       string `yaml:"prefix"`         // 日志前缀
	Director     string `yaml:"director"`       // 日志文件存储目录
	ShowLine     bool   `yaml:"show_line"`      // 是否显示代码所在行号
	LogInConsole bool   `yaml:"log_in_console"` // 是否在控制台输出日志
}
