package common

import (
	"blog/global"
	"blog/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//var list []models.BannerModel
	DB := global.DB
	if option.Debug {
		DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	count = global.DB.Select("id").Find(&list).RowsAffected
	if option.Page == 0 {
		option.Page = 1
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
