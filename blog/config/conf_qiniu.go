package config

type Qiniu struct {
	Enable    string  `yaml:"enable" json:"enable"`
	AccessKey string  `yaml:"accessKey" json:"accessKey"`
	SecretKey string  `yaml:"secretKey" json:"secretKey"`
	Bucket    string  `yaml:"bucket" json:"bucket"` //存储桶名字
	CON       string  `yaml:"CON" json:"CON"`       //访问图片地址前缀
	Zone      string  `yaml:"zone" json:"zone"`     //存储地区
	Size      float64 `yaml:"size" json:"size"`     //存储大小限制 mb
}
