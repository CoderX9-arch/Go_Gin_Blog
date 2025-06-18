package img_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

type ImgResponse struct {
	ID   uint64 `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImgApi) ImgNameListView(c *gin.Context) {
	//var cr models.PageInfo
	//err := c.ShouldBindJSON(&cr)
	//if err != nil {
	//	res.FailWithCode(res.ArgumentError, c)
	//	return
	//}
	var imgeList []ImgResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imgeList)
	//list, count, err := common.ComList(models.BannerModel{}, common.Option{
	//	cr,
	//	true,
	//})
	//res.OkWithList(count, list, c)
	res.OkWithData(imgeList, c)
}
