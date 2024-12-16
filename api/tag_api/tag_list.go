package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})

	if err != nil {
		res.FailWithMessage("查询标签失败", c)
		return
	}
	//需要展示这个标签文章的数量
	res.OkWithList(list, count, c)
}
