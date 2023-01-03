package model

import "time"

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

type User struct {
	UserID   int       `db:"uid"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	Email    string    `db:"email"`
	CreateAt time.Time `db:"create_time"`
	UpdateAt time.Time `db:"update_time"`
}

type ApiUser struct {
	Token string `json:"token"`
}
