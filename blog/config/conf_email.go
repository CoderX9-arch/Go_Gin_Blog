package config

type Email struct {
	Host            string `json:"host" yaml:"host"`
	Port            int    `json:"port" yaml:"port"`
	User            string `json:"user" yaml:"user"`
	Password        string `json:"password" yaml:"password"`
	DefaltFormEmail string `json:"defalt_form_email" yaml:"defalt_form_email"` //默认发件人
	UseSSL          bool   `json:"use_ssl" yaml:"use_ssl"`
	UserTls         bool   `json:"user_tls" yaml:"user_tls"`
}
