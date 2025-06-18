package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`            // 更新时间
}

type PageInfo struct {
	Page  int    `form:"page"` //当前页
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type DELRequest struct {
	IDList []uint `json:"id_list"`
}
