package config

import "fmt"

type System struct {
	Host string `yaml:"host"` // 系统主机地址
	Port int    `yaml:"port"` // 系统端口号
	Env  string `yaml:"env"`  // 系统环境
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
