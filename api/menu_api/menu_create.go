package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title        string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path         string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan       string      `json:"slogan" structs:"slogan"`
	Abstract     ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime int         `json:"abstract_time" structs:"abstract_time"`
	// 切换的时间，单位秒
	BannerTime int `json:"banner_time" structs:"banner_time"`
	// 切换的时间，单位秒
	Sort int `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	// 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
	// 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复值的判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.OkWithMessage("菜单数据已存在", c)
		return
	}
	//创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err := global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}

	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel

	for _, sort := range cr.ImageSortList {
		//这里也要判断image——id是否真的有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}

	//给第三张表入库
	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
