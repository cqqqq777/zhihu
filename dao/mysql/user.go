package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	FindUserByUsername = "select count(uid) from users where username = ?"
	FindUserByEmail    = "select count(uid) from users where email = ?"
	AddUSER            = "insert into users(uid,username,password,email) values(?,?,?,?)"
)

func CheckUsername(username string) error {
	var count int64
	if err := g.Mdb.Get(&count, FindUserByUsername, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func CheckEmail(email string) error {
	var count int64
	if err := g.Mdb.Get(&count, FindUserByEmail, email); err != nil {
		return err
	}
	if count > 0 {
		return ErrorEmailExist
	}
	return nil
}

func AddUser(user *model.User) error {
	_, err := g.Mdb.Exec(AddUSER, user.UserID, user.Username, user.Password, user.Email)
	return err
}
