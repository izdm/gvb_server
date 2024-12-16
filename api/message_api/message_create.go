package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容
}

// 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	//应该是当前用户发布消息  sendid就是当前登陆人的id
	var cr MessageRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &err, c)
		return
	}
	var sendUser, recvUser models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:     cr.SendUserID,
		SendUserName:   sendUser.UserName,
		SendUserAvatar: sendUser.Avatar,
		RevUserID:      cr.RevUserID,
		RevUserName:    recvUser.NickName,
		RevUserAvatar:  recvUser.Avatar,
		IsRead:         false,
		Content:        cr.Content,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
}
