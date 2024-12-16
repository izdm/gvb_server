package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
}
