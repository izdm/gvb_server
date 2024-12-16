package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询用户id"` // 查询的目标用户 ID
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	// 解析请求体
	var cr MessageRecordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c) // 使用统一错误处理逻辑
		return
	}

	// 获取当前登录用户信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 查询与目标用户的所有消息记录
	var messageList []models.MessageModel
	err := global.DB.Order("created_at asc").Find(&messageList,
		"(send_user_id = ? AND rev_user_id = ?) OR (send_user_id = ? AND rev_user_id = ?)",
		claims.UserID, cr.UserID, cr.UserID, claims.UserID).Error

	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("获取消息记录失败", c)
		return
	}

	// 返回查询结果
	res.OkWithData(messageList, c)
}
