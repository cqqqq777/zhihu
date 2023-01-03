package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code RespCode    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Resp{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func RespFailed(c *gin.Context, code RespCode) {
	c.JSON(http.StatusOK, &Resp{
		Code: code,
		Msg:  code.Msg(),
	})
}
