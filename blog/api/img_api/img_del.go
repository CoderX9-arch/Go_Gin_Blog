package img_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImgApi) ImgDelView(c *gin.Context) {
	var cr models.DELRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imgList []models.BannerModel
	count := global.DB.Find(&imgList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("该图片不存在", c)
		return
	}
	global.DB.Delete(&imgList, cr.IDList)
	res.FailWithMsg(fmt.Sprintf("成功删除图片%d", count), c)
}
