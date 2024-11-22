package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/", settingsApi.SettingsInfoView)
	router.GET("/settingstest", settingsApi.SettingsInfoViewTest) //测试respose函数
	router.GET("/sTest2", settingsApi.SettingsInfoViewTest2)
}
