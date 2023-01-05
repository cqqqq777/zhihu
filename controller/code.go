package controller

type RespCode int16

const (
	CodeSuccess      RespCode = 0
	CodeInvalidParam RespCode = 1000 + iota
	CodeUserExist
	CodeEmailExist
	CodeUserNotExist
	CodeEmailNotExist
	CodeWrongPassword
	CodeWrongVerification
	CodeServiceBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidId
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeEmailExist:        "邮箱已注册",
	CodeUserNotExist:      "用户不存在",
	CodeEmailNotExist:     "邮箱未注册",
	CodeWrongPassword:     "密码错误",
	CodeWrongVerification: "验证码错误",
	CodeServiceBusy:       "服务器繁忙",
	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的token",
	CodeInvalidId:         "无效的id",
}

func (code RespCode) Msg() string {
	return codeMsgMap[code]
}
