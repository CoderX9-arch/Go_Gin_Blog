package settings_api

import (
	"blog/config"
	"blog/core"
	"blog/global"
	"blog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailUpdateView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println("before:", global.Config)
	global.Config.Email = cr
	fmt.Println("after:", global.Config)

	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWith(c)
}
