package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "zhihu/global"
	"zhihu/services"
)

func GetAllTopic(c *gin.Context) {
	data, err := services.GetAllTopic()
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		g.Logger.Error(fmt.Sprintf("%v", err))
		return
	}
	RespSuccess(c, data)
}
