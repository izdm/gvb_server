package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.GET("/image_names", app.ImageNameListView)
	router.POST("/images", app.ImageUploadView)
	router.GET("/images", app.ImageListView)
	router.DELETE("/images", app.ImageRemoveView)
	router.PUT("/images", app.ImageUpdateView)
}
