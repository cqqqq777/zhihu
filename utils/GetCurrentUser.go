package utils

import (
	"github.com/gin-gonic/gin"
	"zhihu/model"
)

func GetCurrentUser(c *gin.Context) (UserID int, ok bool) {
	uid, ok := c.Get(model.CtxGetUID)
	if !ok {
		return
	}
	UserID, ok = uid.(int)
	if !ok {
		return
	}
	return
}
