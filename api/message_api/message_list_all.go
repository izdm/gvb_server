package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
	}

	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})

	res.OkWithList(list, count, c)
}
