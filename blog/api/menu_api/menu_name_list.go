package menu_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

type MenuNameListResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameListResponse
	//c.ShouldBindJSON(&menuNameList)
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
