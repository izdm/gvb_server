package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (UserApi) UserRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//TODO：删除用户 消息表 评论表 用户收藏的文章 用户发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)

}
