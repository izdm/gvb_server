package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	if err := c.ShouldBind(&cr); err != nil {
		//参数出了问题
		res.FailWithError(err, &cr, c)
		return
	}

	//重复的判断 每次添加之前去查一下有没有重复的title
	var tag models.TagModel
	err := global.DB.Take(&tag, "id = ?", id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}
