package user_ser

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/pwd"
)

const Avatar = "/uploads/avatar/default.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name", userName).Error
	if err == nil {
		//说明用户名存在
		global.Log.Error("用户名已存在，请重新输入")
	}

	//头像问题
	//1.默认头像

	//对密码进行hash
	hashPwd := pwd.HashPwd(password)
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       "内网地址",
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		return err
	}
	return nil
}
