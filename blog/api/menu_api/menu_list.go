package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type menuListResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	var menulist []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menulist).Select("id").Scan(&menuIDList)
	fmt.Println(menulist, menuIDList)

	var menuImageModel []models.MenuImageModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuImageModel, "menu_id in?", menuIDList)
	var menuList []menuListResponse
	for _, model := range menulist {
		banners := []Banner{}
		for _, banner := range menuImageModel {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.MenuID,
				Path: banner.BannerModel.Path,
			})
		}
		menuList = append(menuList, menuListResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menuList, c)
	return
}
