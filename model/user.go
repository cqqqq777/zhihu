package model

import (
	"time"
)

const CtxGetUID = "UserID"

type ParamRegisterUser struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re-password" binding:"required"`
	//验证码
	Verification int64 `json:"verification" binding:"required"`
}

type ParamLoginUser struct {
	UsernameOrEmail string `json:"username/email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type ParamReviseUser struct {
	Uid         int
	NewUsername string `json:"new-username"`
	OriPassword string `json:"ori-password"`
	NewPassword string `json:"new-password"`
	RePassword  string `json:"re-password"`
}

type User struct {
	Id           int       `json:"id,omitempty" db:"id"`
	Uid          int       `json:"uid" db:"uid"`
	Gender       int       `json:"gender" db:"gender"`
	Introduction string    `json:"introduction" db:"introduction"`
	Username     string    `json:"username" db:"username"`
	Password     string    `json:"-" db:"password"`
	Email        string    `json:"email" db:"email"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}

type ApiUser struct {
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}
