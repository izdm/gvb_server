package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo false "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	referer := c.Request.Header.Get("Referer")
	isAdmin := strings.Contains(referer, "/admin")

	var count int64

	// 如果是 admin 请求，展示所有广告，否则仅展示 is_show 为 true 的广告
	query := models.AdvertModel{}
	if !isAdmin {
		query.IsShow = true
	}

	list, count, err := common.ComList(query, common.Option{
		PageInfo: cr,
	})

	if err != nil {
		res.FailWithMessage("查询广告列表失败", c)
		return
	}

	res.OkWithList(list, count, c)
}
