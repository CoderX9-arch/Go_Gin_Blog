package settings_api

import (
	"blog/config"
	"blog/core"
	"blog/global"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//fmt.Println(cr.Name)
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		fmt.Println(info)
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		fmt.Println(info)
		global.Config.SiteInfo = info

	case "email":
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = email
	case "qq":
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = qq
	case "qiniu":
		var qiniu config.Qiniu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Qiniu = qiniu
	case "jwt":
		var jwt config.Jwy
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwy = jwt
	default:
		res.FailWithMsg("没有对应信息", c)
		return
	}
	core.SetYaml()
	res.OkWith(c)
}
