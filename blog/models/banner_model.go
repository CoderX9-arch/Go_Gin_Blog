package models

import (
	"blog/global"
	"blog/models/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片hash值，判断重复图片
	Name string `gorm:"size:38" json:"name"` //图片名称
	//ArticleModels []ArticleModel `gorm:"foreigKey：CoverID" json:"-"`
	ImgType ctype.ImgType `gorm:"default:1" json:"img_type"` //图片类型：本地/云
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImgType == ctype.Local {
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return
}
