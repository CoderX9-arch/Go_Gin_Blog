package models

// MenuImageModel 菜单图片关联表
type MenuImageModel struct {
	MenuID      uint        `json:"menu_id"`
	BannerID    uint        `json:"banner_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID" json:"menu_model"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID" json:"image_model"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
