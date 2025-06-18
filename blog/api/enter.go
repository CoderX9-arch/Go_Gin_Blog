package api

import (
	"blog/api/img_api"
	"blog/api/menu_api"
	"blog/api/settings_api"
	"blog/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImgApi      img_api.ImgApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
