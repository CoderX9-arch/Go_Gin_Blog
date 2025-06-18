package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	Redis    Redis    `yaml:"redis"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Qiniu    Qiniu    `yaml:"qiniu"`
	Email    Email    `yaml:"email"`
	Jwy      Jwy      `yaml:"jwy"`
	Upload   Upload   `yaml:"upload"`
}
