package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	FindUserByUsername = "select * from users where username = ?"
	FindUserByEmail    = "select * from users where email = ?"
)

func CheckUsername(username string) bool {
	user := new(model.User)
	if err := g.Mdb.Get(user, FindUserByUsername); err != nil {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	user := new(model.User)
	if err := g.Mdb.Get(user, FindUserByEmail); err != nil {
		return false
	}
	return true
}
