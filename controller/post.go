package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
