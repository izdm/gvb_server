package config

import "fmt"

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	Upload   Upload   `yaml:"upload"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	Redis    Redis    `yaml:"redis"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
