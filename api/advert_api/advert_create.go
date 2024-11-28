package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// binding：用于字段的验证规则，常见于 Gin 等 Web 框架。在这里，binding 标签用于指定字段的验证规则，比如 required（必填）
// 、url（URL 格式），等。Gin 会根据这些规则对请求的参数进行验证。如果参数不符合规则，会触发验证错误。
// msg：这是自定义的错误消息标签，用来为字段验证失败时提供具体的错误信息。通常，框架或库（如 Gin）会将验证错误消息显示给客户端，msg 标签让你可以定制这些消息。
type AdvertModel struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片链接非法" structs:"images"` // 图片
	IsShow bool   `json:"is_show" msg:"请选择是否展示" structs:"is_show"`                     // 是否展示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertModel true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertModel
	if err := c.ShouldBind(&cr); err != nil {
		//参数出了问题
		res.FailWithError(err, &cr, c)
		return
	}

	//重复的判断 每次添加之前去查一下有没有重复的title
	var advert models.AdvertModel
	err := global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("广告已存在", c)
		return
	}

	if err := global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)
}
