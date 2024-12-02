package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type UserRole struct {
	Role   ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	UserID uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户ID错误，用户不存在", c)
	}
	err = global.DB.Model(&user).Update("role", cr.Role).Error
	if err != nil {
		res.FailWithMessage("权限修改失败", c)
		return
	}
	res.OkWithMessage("权限修改成功", c)
}
