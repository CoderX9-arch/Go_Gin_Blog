package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/ctype"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImgSort struct {
	ImgID uint `json:"img_id"`
	Sort  int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请输入菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	ImageSortList []ImgSort   `json:"image_sort_list" `
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithErr(err, &cr, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		global.Log.Error(err)
		res.FailWithMsg("菜单重复", c)
		return
	}
	menuModel := models.MenuModel{
		Title:    cr.Title,
		Path:     cr.Path,
		Slogan:   cr.Slogan,
		Abstract: cr.Abstract,
		MenuTime: cr.BannerTime,
		Sort:     cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMsg("菜单添加成功", c)
		return
	}

	var menu_image_models []models.MenuImageModel
	for _, sort := range cr.ImageSortList {
		fmt.Println(sort.ImgID)
		menu_image_models = append(menu_image_models, models.MenuImageModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImgID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Debug().Create(&menu_image_models).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	res.OkWithMsg("菜单添加成功", c)
}
