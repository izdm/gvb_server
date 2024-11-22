package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //设置运行时模式
	router := gin.Default()

	//路由分组
	apiRouterGroup := router.Group("api")
	RouterGroupApp := RouterGroup{apiRouterGroup}

	//路由分层
	//系统配置api
	RouterGroupApp.SettingsRouter()
	return router
}
