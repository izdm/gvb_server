package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/Email"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	//用户绑定邮箱，第一次输入是邮箱
	//后台会给这个邮箱发验证码
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr BindEmailRequest

	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次，后台发验证码
		//生成四位验证码，将生成的验证码存入session
		code := random.Code(4)
		//写入session
		session.Set("valid_code", code)
		session.Set("email", cr.Email)

		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = Email.NewCode().Send(cr.Email, "你的验证码是:"+code)
		if err != nil {
			global.Log.Error(err)
			return
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}

	// 第二次请求时，进行验证
	storedEmail := session.Get("email") // 获取第一次存储的邮箱
	if storedEmail == nil {
		res.FailWithMessage("未找到绑定邮箱信息", c)
		return
	}

	// 校验输入的邮箱和第一次的邮箱是否一致
	if cr.Email != storedEmail {
		res.FailWithMessage("邮箱与第一次输入的邮箱不一致", c)
		return
	}

	//第二次 用输入邮箱与验证码 密码
	code := session.Get("valid_code")
	//校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	//修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	//第一次的邮箱和第二次的邮箱也要做一次校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		res.FailWithMessage("绑定邮箱失败", c)
	}
	res.OkWithMessage("绑定邮箱成功", c)
	//用户完成绑定
}
