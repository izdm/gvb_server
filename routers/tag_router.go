package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tag", app.TagCreateView)
	router.GET("tag", app.TagListView)
	router.PUT("tag/:id", app.TagUpdateView)
	router.DELETE("tag", app.TagRemoveView)
}
