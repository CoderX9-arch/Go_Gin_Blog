package routers

import "blog/api"

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menu", app.MenuCreateView)
	router.GET("menu", app.MenuListView)
	router.GET("menunames", app.MenuNameListView)
	router.PUT("menu_update/:id", app.MenuUpdateView)
	router.DELETE("menu_del", app.MenuDelView)
	router.GET("menu/:id", app.MenuDetailView)
}
