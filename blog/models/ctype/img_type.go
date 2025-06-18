package ctype

import "encoding/json"

type ImgType int

const (
	Local ImgType = 1 //本地
	QiNiu ImgType = 2 //云
)

func (s ImgType) MarshJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImgType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"

	}
	return str
}
