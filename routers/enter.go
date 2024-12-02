package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gvb_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //设置运行时模式
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//路由分组
	//将路由分组，所有添加到 apiRouterGroup 的路由都将以 /api 开头。
	//例如，apiRouterGroup.GET("/adverts", handler) 实际路径是 /api/adverts。
	apiRouterGroup := router.Group("api") //创建一个路由组
	RouterGroupApp := RouterGroup{apiRouterGroup}

	//路由分层
	//系统配置api
	RouterGroupApp.SettingsRouter()
	RouterGroupApp.ImagesRouter()
	RouterGroupApp.AdvertRouter()
	RouterGroupApp.MenuRouter()
	RouterGroupApp.UserRouter()
	return router
}
