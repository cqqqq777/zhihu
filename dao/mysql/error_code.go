package mysql

import "errors"

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("用户不存在")
	ErrorEmailExist    = errors.New("邮箱已注册")
	ErrorEmailNotExist = errors.New("邮箱未注册")
)
