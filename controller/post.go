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
	if post.Type != 1 && post.Type != 2 {
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
	//获取路径参数并校验
	pidStr := c.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//获取数据
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

// QuestionList 分页获取问题列表
func QuestionList(c *gin.Context) {
	//获取分页参数
	page, size := utils.GetPageInfo(c)
	data, err := services.QuestionList(page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

// EssayList 分页获取文章列表
func EssayList(c *gin.Context) {
	//获取分页参数
	page, size := utils.GetPageInfo(c)
	data, err := services.EssayList(page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

// UserQuestionList 获取用户提出的所有问题
func UserQuestionList(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.UserQuestionList(page, size, int64(uid))
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

// UserEssayList 获取用户发表的所有文章
func UserEssayList(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.UserEssayList(page, size, int64(uid))
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}
