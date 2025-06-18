package settings_api

import (
	"blog/global"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	//var cr config.SiteInfo
	//err := c.ShouldBindJSON()
	//res.Ok(map[string]string{}, "xxx", c)
	res.OkWithData(global.Config.Email, c)

}
