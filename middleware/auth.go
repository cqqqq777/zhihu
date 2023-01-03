package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zhihu/controller"
	"zhihu/model"
	"zhihu/utils"
)

func JWTAuth(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	//假设Token放在Header的Authorization中，并使用Bearer开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		controller.RespFailed(c, controller.CodeNeedLogin)
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		controller.RespFailed(c, controller.CodeInvalidToken)
		c.Abort()
		return
	}
	myClaim, err := utils.ParseToken(parts[1])
	if err != nil {
		controller.RespFailed(c, controller.CodeInvalidToken)
		c.Abort()
		return
	}
	c.Set(model.CtxGetUID, myClaim.Uid)
}
