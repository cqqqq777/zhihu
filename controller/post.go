package controller

import (
	"errors"
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
	uidStr := c.Request.Header.Get("uid")
	var uid int64
	if uidStr == "" {
		uid = 0
	} else {
		uid, err = strconv.ParseInt(uidStr, 10, 64)
	}
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := services.PostDetail(pid, uid)
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

// UpdatePost 更新一篇帖子
func UpdatePost(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	pidStr := c.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	post := new(model.Post)
	post.AuthorID = uid
	post.Pid = int(pid)
	err = c.ShouldBindJSON(post)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if post.Type != 1 && post.Type != 2 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	err = services.UpdatePost(post)
	if err != nil {
		if errors.Is(err, mysql.ErrorNoPermission) {
			RespFailed(c, CodeNoPermission)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, nil)
}

// DeletePost 删除某一条帖子
func DeletePost(c *gin.Context) {
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	pidStr := c.Param("pid")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	err = services.DeletePost(uid, int(pid))
	if err != nil {
		if errors.Is(err, mysql.ErrorNoPermission) {
			RespFailed(c, CodeNoPermission)
			return
		}
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, nil)
}

func SearchPost(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	page, size := utils.GetPageInfo(c)
	data, err := services.SearchPosts(page, size, key)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}

func GetHotPostList(c *gin.Context) {
	page, size := utils.GetPageInfo(c)
	data, err := services.GetHotPostList(page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		return
	}
	RespSuccess(c, data)
}

func RecommendPost(c *gin.Context) {
	data, err := services.RecommendPost()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(err.Error())
		return
	}
	RespSuccess(c, data)
}
