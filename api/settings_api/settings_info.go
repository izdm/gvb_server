package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// 显示某一项的配置信息，这样子把多个接口合并成一个
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)

	}

}

//func (SettingsApi) SettingsInfoViewTest(c *gin.Context) {
//	res.OkWithData(map[string]string{
//		"id": "zdm",
//	}, c)
//}
//func (SettingsApi) SettingsInfoViewTest2(c *gin.Context) {
//	res.FailWithCode(res.SettingsError, c)
//}
