package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao/mysql"
	g "zhihu/global"
	"zhihu/services"
)

// GetAllTopic 获取全部话题
func GetAllTopic(c *gin.Context) {
	data, err := services.GetAllTopic()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

// TopicDetail 获取话题详情
func TopicDetail(c *gin.Context) {
	tidStr := c.Param("tid")
	if tidStr == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	data, err := services.TopicDetail(tid)
	if err != nil {
		if err == mysql.ErrorInvalidId {
			RespFailed(c, CodeInvalidId)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}
