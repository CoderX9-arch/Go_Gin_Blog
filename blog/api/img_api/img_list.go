package img_api

import (
	"blog/models"
	"blog/models/res"
	"blog/service/common"
	"github.com/gin-gonic/gin"
)

func (ImgApi) ImgListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		cr,
		true,
	})
	res.OkWithList(count, list, c)
	return
}
