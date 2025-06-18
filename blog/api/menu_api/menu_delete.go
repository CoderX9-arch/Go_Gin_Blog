package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuDelView(c *gin.Context) {
	var cr models.DELRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("该菜单不存在", c)
		return
	}
	//global.DB.Model(&menuList).Association("BannerModel").Clear()
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Delete(&menuList, cr.IDList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除菜单失败", c)
		return
	}

	res.FailWithMsg(fmt.Sprintf("成功删除菜单%d", count), c)
}
