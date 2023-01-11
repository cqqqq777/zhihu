package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao/mysql"
	g "zhihu/global"
	"zhihu/model"
	"zhihu/services"
	"zhihu/utils"
)

// CommentPost 给帖子评论
func CommentPost(c *gin.Context) {
	comment := new(model.Comment)
	err := c.ShouldBindJSON(comment)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, ok := utils.GetCurrentUser(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	comment.AuthorId = uid
	comment.Cid, err = utils.GetID()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	commentedUid, err := mysql.GetAuthorIdByPid(int(comment.PostId))
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	comment.CommentedUid = int64(commentedUid)
	err = services.CommentPost(comment)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, nil)
}

func PostCommentList(c *gin.Context) {
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
	page, size := utils.GetPageInfo(c)
	data, err := services.PostCommentList(pid, uid, page, size)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Warn(err.Error())
		return
	}
	RespSuccess(c, data)
}
