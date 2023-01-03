package controller

import (
	"github.com/gin-gonic/gin"
	"zhihu/model"
	"zhihu/services"
)

// PostVerification 发送验证码
func PostVerification(c *gin.Context) {
	//获取邮箱并校验
	email := c.PostForm("email")
	if email == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//给邮箱发验证码
	if err := services.PostVerification(email); err != nil {
		RespFailed(c, CodeServiceBusy)
		return
	}
	//返回响应
	RespSuccess(c, nil)
}

// Register 注册
func Register(c *gin.Context) {
	ParamUser := new(model.ParamUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.Username == "" || ParamUser.Password == "" || ParamUser.RePassword == "" || ParamUser.Email == "" || ParamUser.Password != ParamUser.RePassword {
		RespFailed(c, CodeInvalidParam)
		return
	}

}
