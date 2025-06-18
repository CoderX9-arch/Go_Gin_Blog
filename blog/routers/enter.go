package routers

import (
	"blog/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	routerGroup := RouterGroup{router}
	routerGroup.InitSettingsRouter()
	routerGroup.ImgRouter() //图片上传
	routerGroup.MenuRouter()
	routerGroup.UserRouter()
	return router
}
