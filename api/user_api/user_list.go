package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desense"
	"gvb_server/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	//判断是不是超级管理员
	// 获取 claims
	_claims, exists := c.Get("claims")
	if !exists || _claims == nil {
		res.FailWithMessage("未拿到数据", c) // 如果 claims 不存在，则返回 401 错误
		return
	}

	// 类型断言
	claims, ok := _claims.(*jwts.CustomClaims)
	if !ok {
		res.FailWithMessage("断言失败", c) // 如果类型断言失败，则返回 401 错误
		return
	}
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	var users []models.UserModel
	for _, user := range list {

		fmt.Println(user.Role)
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//非管理员
			user.UserName = ""
		}
		//脱敏
		user.Email = desense.DesensitizationEmail(user.Email)
		user.Tel = desense.DesensitizationTel(user.Tel)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
