package routers

import "blog/api"

func (router RouterGroup) ImgRouter() {
	app := api.ApiGroupApp.ImgApi
	router.POST("image", app.ImgUploadView)
	router.GET("image", app.ImgListView)
	router.DELETE("image", app.ImgDelView)
	router.PUT("image", app.ImguUPDNameView)
	router.GET("image_names", app.ImgNameListView)
}
