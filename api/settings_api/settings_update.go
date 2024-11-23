package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	//global.Log.Info("111")
	//global.Log.Info(cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error(), c)
		return
	}
	res.OkWith(c)
}
