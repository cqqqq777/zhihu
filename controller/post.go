package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao/mysql"
	g "zhihu/global"
	"zhihu/model"
	"zhihu/services"
	"zhihu/utils"
)

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	post := new(model.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if post.Type != 0 && post.Type != 1 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	post.AuthorID = uid
	if err := services.CreatePost(post); err != nil {
		if err == mysql.ErrorQuestionExist {
			RespFailed(c, CodeQuestionExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, nil)
}

// PostDetail 获取帖子详情
func PostDetail(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	data, err := services.PostDetail(pid)
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
