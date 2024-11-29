package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	//router.GET("/image_names", app.ImageNameListView)
	router.POST("/menus", app.MenuCreateView)
	router.GET("/menus", app.MenuListView)
	router.GET("/menu_names", app.MenuNameList)
	//router.GET("/images", app.ImageListView)
	//router.DELETE("/images", app.ImageRemoveView)
	//.PUT("/images", app.ImageUpdateView)
}
