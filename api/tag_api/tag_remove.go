package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	var TagList []models.TagModel
	result := global.DB.Find(&TagList, cr.IDList)
	count := result.RowsAffected
	if count < 1 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	//如果这个标签下有文章怎么办？

	global.DB.Delete(&TagList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
