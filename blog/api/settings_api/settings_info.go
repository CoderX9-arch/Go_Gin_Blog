package settings_api

import (
	"blog/global"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	//var cr config.SiteInfo
	//err := c.ShouldBindJSON()
	//res.Ok(map[string]string{}, "xxx", c)
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println(cr.Name)
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.Qiniu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwy, c)
	default:
		res.FailWithMsg("没有对应信息", c)
	}
	//res.OkWithData(global.Config.SiteInfo, c)

}
