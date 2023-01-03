package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	g "zhihu/global"
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
		g.Logger.Debug(fmt.Sprintf("%v", err))
		return
	}
	//返回响应
	RespSuccess(c, nil)
}

// Register 注册
func Register(c *gin.Context) {
	//获取参数并校验
	ParamUser := new(model.ParamUser)
	if err := c.ShouldBindJSON(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.Username == "" || ParamUser.Password == "" || ParamUser.RePassword == "" || ParamUser.Email == "" || ParamUser.Password != ParamUser.RePassword {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//根据错误类型返回响应
	if err := services.Register(ParamUser); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			RespFailed(c, CodeUserExist)
			return
		}
		if errors.Is(err, mysql.ErrorEmailExist) {
			RespFailed(c, CodeEmailExist)
			return
		}
		if errors.Is(err, redisdao.ErrorInvalidVerification) {
			RespFailed(c, CodeWrongVerification)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Debug(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, nil)
}
