package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithErr(err, &cr, c)
		return
	}
	id := c.Param("id")

	var menuModel models.MenuModel
	err = global.DB.Debug().Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuImageModel
		for _, v := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuImageModel{
				MenuID:   menuModel.ID,
				BannerID: v.ImgID,
				Sort:     menuModel.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			res.FailWithMsg("创建菜单失败", c)
			return
		}
	}

	//err = global.DB.Model(&menuModel).Updates(m2).Error
	err = global.DB.Model(&menuModel).Updates(map[string]interface{}{
		"title":     cr.Title,
		"path":      cr.Path,
		"slogan":    cr.Slogan,
		"abstract":  cr.Abstract,
		"menu_time": cr.BannerTime,
		"sort":      cr.Sort}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OkWithMsg("修改菜单成功", c)
}
