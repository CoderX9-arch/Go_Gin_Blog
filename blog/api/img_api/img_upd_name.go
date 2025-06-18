package img_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

type ImgUPDNameRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入图片id"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

func (ImgApi) ImguUPDNameView(c *gin.Context) {
	var cr ImgUPDNameRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithErr(err, &cr, c)
		return
	}
	var ImgModel models.BannerModel
	err = global.DB.Take(&ImgModel, &cr.ID).Error
	if err != nil {
		res.FailWithMsg("图片不存在", c)
		return
	}
	err = global.DB.Model(&ImgModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithMsg("图片名称修改成功", c)
}
