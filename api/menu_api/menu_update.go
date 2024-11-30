package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest

	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	id := c.Param("id")
	//先清空之前的banner
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 更新其他字段
	menuModel.Title = cr.Title
	menuModel.Path = cr.Path
	menuModel.Slogan = cr.Slogan
	menuModel.Abstract = cr.Abstract
	menuModel.AbstractTime = cr.AbstractTime
	menuModel.BannerTime = cr.BannerTime
	menuModel.Sort = cr.Sort

	err = global.DB.Save(&menuModel).Error
	if err != nil {
		res.FailWithMessage("更新菜单信息失败", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选择了banner，那就去添加
	if len(cr.ImageSortList) > 0 {
		//操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}

		err = global.DB.Create(&bannerList).Error

		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	res.OkWithMessage("修改菜单成功", c)
}
