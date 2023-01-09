package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	RespSuccess(c, nil)
}
