package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// 处理分页
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBind(&cr)
	fmt.Println(cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
