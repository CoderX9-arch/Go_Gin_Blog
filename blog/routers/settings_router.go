package routers

import (
	"blog/api"
)

func (r RouterGroup) InitSettingsRouter() {
	settingapi := api.ApiGroupApp.SettingsApi
	r.GET("settings/:name", settingapi.SettingsInfoView)
	r.PUT("settings/:name", settingapi.SettingsUpdateView)
	//r.GET("settings_email", settingapi.SettingsEmailInfoView)
	//r.PUT("settings_email", settingapi.SettingsEmailUpdateView)
	//app := api.ApiGroupApp.ImgApi
	//r.POST("image", app.ImgUploadView)
}
