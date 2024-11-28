package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertRemoveView 广告批量删除
// @Tags 广告管理
// @Summary 广告批量删除
// @Description 广告批量删除
// @Param data body models.PageInfo true "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	var advertList []models.AdvertModel
	result := global.DB.Find(&advertList, cr.IDList)
	count := result.RowsAffected
	if count < 1 {
		res.FailWithMessage("广告不存在", c)
		return
	}

	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)
}
