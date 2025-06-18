package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel models.MenuModel
	//var menuIDList []uint
	var count int64
	err := global.DB.Find(&menuModel, id).Count(&count).Error
	fmt.Println(err)
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	if err != nil {
		global.Log.Error(err)
		return
	}

	var menuImageModel []models.MenuImageModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuImageModel, "menu_id in (?)", id)
	//var menuResponse menuListResponse

	banners := []Banner{}
	for _, banner := range menuImageModel {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.MenuID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := menuListResponse{
		menuModel,
		banners,
	}
	res.OkWithData(menuResponse, c)
	return
}
