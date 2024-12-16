package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"time"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`    // 消息内容
	CreatedAt        time.Time `json:"created_at"` // 创建时间
	MessageCount     int       `json:"message_count"`
}

type MessageGroup map[string]*Message

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message

	// 查找当前用户的所有相关消息
	global.DB.Order("created_at").Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)

	for _, model := range messageList {
		// 生成唯一的分组标识符
		idNum := fmt.Sprintf("%d-%d",
			min(model.SendUserID, model.RevUserID),
			max(model.SendUserID, model.RevUserID),
		)

		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}

		// 检查是否已有该组的消息
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &message
		} else {
			val.MessageCount += 1
			// 更新最近一条消息内容
			val.Content = model.Content
			val.CreatedAt = model.CreatedAt
		}
	}

	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)
	return
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
