package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	if err := c.ShouldBind(&cr); err != nil {
		//参数出了问题
		res.FailWithError(err, &cr, c)
		return
	}

	//重复的判断 每次添加之前去查一下有没有重复的title
	var tag models.TagModel
	err := global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("标签已存在", c)
		return
	}

	if err := global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)
}
