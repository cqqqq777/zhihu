package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao/redisdao"
	g "zhihu/global"
	"zhihu/services"
	"zhihu/utils"
)

func StarPost(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	err = services.StarPost(pid, int64(uid))
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	redisdao.ClearPostCache(pid)
	RespSuccess(c, nil)
}

func StarComment(c *gin.Context) {
	cidStr := c.Param("cid")
	cid, err := strconv.ParseInt(cidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	if err = services.StarComment(cid, int64(uid)); err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}
